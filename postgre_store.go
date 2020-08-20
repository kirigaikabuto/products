package products

import (
	"database/sql"
	"errors"
	"log"
	_ "github.com/lib/pq"
)

var Queries = []string{
	`CREATE TABLE IF NOT EXISTS products (
		id serial,
		name text,
		price int,
		image_url text,
		PRIMARY KEY(id)
	);`,
}

type postgreStore struct {
	db *sql.DB
}

func NewPostgreStore(cfg Config) (ProductStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgreStore{db: db}, err
}

func (ps *postgreStore) List() ([]Product, error) {
	var products []Product
	data, err := ps.db.Query("select id,name,price,image_url from products")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		product := Product{}
		err = data.Scan(&product.Id,&product.Name,&product.Price,&product.ImageUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (ps *postgreStore) Create(product *Product) (*Product, error) {
	result, err := ps.db.Exec("insert into products (name,price,image_url) values ($1,$2,$3) RETURNING id",product.Name,product.Price,product.ImageUrl)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("not defined error")
	}
	return product, nil
}

func (ps *postgreStore) GetById(id int64) (*Product, error) {
	product := &Product{}
	return product, nil
}

func (ps *postgreStore) Update(product *Product) (*Product, error) {
	return product, nil
}

func (ps *postgreStore) Delete(id int64) error {
	return nil
}
