package repository

import (
	"github.com/stretchr/testify/mock"
	"unit_test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := repo.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*entity.Category)
}
