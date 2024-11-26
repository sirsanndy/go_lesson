package repository

import "unit_test/entity"

type CategoryRepository interface {
	FindById(is string) *entity.Category
}
