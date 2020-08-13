package products

type Product struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageUrl string `json:"image_url"`
}

type ProductStore interface {
	List() ([]Product, error)
	Create(product *Product) (*Product, error)
	GetById(id int64) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id int64) error
}
