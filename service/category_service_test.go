package service

import (
	"belajar/unit/entity"
	"belajar/unit/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}
var categoryService = CategoryService{
	Repository: categoryRepository,
}

func TestCategoryService_Get(t *testing.T) {

	//check in database. in this case, the result is nil. so the data is not exist
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	//call service for result
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Mohammad Nowaf",
	}

	//check in database. in this case, the result is nil. so the data is not exist
	categoryRepository.Mock.On("FindById", "2").Return(category)

	//call service for result
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id, "Id should be equal")
	assert.Equal(t, category.Name, result.Name, "Name should be equal")

}
