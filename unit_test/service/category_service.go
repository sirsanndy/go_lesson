package service

import (
	"errors"
	"unit_test/entity"
	"unit_test/repository"
)

type CategoryService interface {
	Get(id string) (*entity.Category, error)
}

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{Repository: repo}
}

func (service *CategoryServiceImpl) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}
