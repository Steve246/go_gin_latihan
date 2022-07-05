package usecase

import (
	"go_gin_latihan/model"
	"go_gin_latihan/repository"
)

type CreateCategoryUsecase interface {
	CreateCategory(newCategory *model.Category) error
}

type createCategoryUsecase struct {
	repo repository.CategoryRepository
}

func (c *createCategoryUsecase) CreateCategory(newCategory *model.Category) error {

	return c.repo.Add(newCategory)

}

func NewCreateProductUseCase(repo repository.CategoryRepository) CreateCategoryUsecase {
	return &createCategoryUsecase{
		repo: repo,
	}
}
