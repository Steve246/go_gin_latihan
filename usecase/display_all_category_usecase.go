package usecase

import (
	"go_gin_latihan/model"
	"go_gin_latihan/repository"
)

type CreateDisplayCategory interface {
	DisplayAll() ([]model.Category, error)
}

type createDisplayCategory struct {
	repo repository.CategoryRepository
}

func (c *createDisplayCategory) DisplayAll() ([]model.Category, error) {

	return c.repo.Retrieve()

}

func NewCreateDisplayCategory(repo repository.CategoryRepository) CreateDisplayCategory {
	return &createDisplayCategory{
		repo: repo,
	}
}
