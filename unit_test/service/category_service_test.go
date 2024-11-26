package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"unit_test/entity"
	"unit_test/repository"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = NewCategoryService(categoryRepository)

func TestCategoryServiceImpl_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryServiceImpl_GetFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "2").Return(&entity.Category{
		Id:   "2",
		Name: "Sandy",
	})
	category, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
}
