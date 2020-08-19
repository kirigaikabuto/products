package products

type ProductService interface {
	ListProducts() ([]Product, error)
	CreateProduct(cmd *CreateProductCommand) (*Product, error)
	GetProductById(cmd *GetProductByIdCommand) (*Product, error)
	UpdateProduct(cmd *UpdateProductCommand) (*Product, error)
	DeleteProduct(cmd *DeleteProductCommand) error
}

type productService struct {
	productStore ProductStore
}

func NewProductService(productStore ProductStore) ProductService {
	return &productService{productStore: productStore}
}

func (ps *productService) ListProducts() ([]Product, error) {
	products, err := ps.productStore.List()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) CreateProduct(cmd *CreateProductCommand) (*Product, error) {
	product := &Product{
		Name:  cmd.Name,
		Price: cmd.Price,
	}
	newProduct, err := ps.productStore.Create(product)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (ps *productService) GetProductById(cmd *GetProductByIdCommand) (*Product, error) {
	product, err := ps.productStore.GetById(cmd.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *productService) UpdateProduct(cmd *UpdateProductCommand) (*Product, error) {
	updateProduct := &Product{}
	if cmd.Price != nil {
		updateProduct.Price = *cmd.Price
	} else if cmd.Name != nil {
		updateProduct.Name = *cmd.Name
	} else if cmd.ImageUrl != nil {
		updateProduct.ImageUrl = *cmd.ImageUrl
	}
	cmdGetProductById := &GetProductByIdCommand{cmd.Id}
	_, err := ps.GetProductById(cmdGetProductById)
	if err != nil {
		return nil, err
	}
	updatedProduct, err := ps.productStore.Update(updateProduct)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (ps *productService) DeleteProduct(cmd *DeleteProductCommand) error {
	cmdGetProductById := &GetProductByIdCommand{cmd.Id}
	_, err := ps.GetProductById(cmdGetProductById)
	if err != nil {
		return err
	}
	err = ps.productStore.Delete(cmd.Id)
	if err != nil {
		return err
	}
	return nil
}
