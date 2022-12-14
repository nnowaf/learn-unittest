package repository

import "belajar/unit/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
