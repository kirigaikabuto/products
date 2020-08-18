package products

import (
	"database/sql"
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
	return products, nil
}

func (ps *postgreStore) Create(product *Product) (*Product, error) {
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
