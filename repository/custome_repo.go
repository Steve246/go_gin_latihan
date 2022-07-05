package repository

import "go_gin_latihan/model"

type CategoryRepository interface {
	Add(newCategory *model.Category) error
	Retrieve() ([]model.Category, error)
}

type categoryRepository struct {
	categoryDb []model.Category
}

func (c *categoryRepository) Retrieve() ([]model.Category, error) {
	return c.categoryDb, nil
}

func (c *categoryRepository) Add(newCategory *model.Category) error {
	c.categoryDb = append(c.categoryDb, *newCategory)

	return nil

}

func NewCategoryRepository() CategoryRepository {
	repo := new(categoryRepository)
	return repo
}
