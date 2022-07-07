package repository

import "go_gin_latihan/model"

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
}

type productRepository struct {
	productDb []model.Product
}

func (p *productRepository) Retrieve() ([]model.Product, error) {
	return p.productDb, nil
}

func (p *productRepository) Add(newProduct *model.Product) error  {
	p.productDb = append(p.productDb, *newProduct)
	return nil

}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}
