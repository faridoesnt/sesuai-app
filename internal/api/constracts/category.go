package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type CategoryRepository interface {
	FindCategory() (categories []entities.Category, err error)
	FindCategoryById(categoryId string) (category entities.Category, err error)
	InsertCategory(category entities.RequestCategory) (err error)
	UpdateCategory(categoryId string, params entities.RequestCategory) (err error)
	DeleteCategory(categoryId string) (err error)
}

type CategoryService interface {
	GetCategory() (listCategory []response.Category)
	GetCategoryDetail(categoryId string) (category response.Category)
	InsertCategory(category entities.RequestCategory) (err error)
	UpdateCategory(categoryId string, params entities.RequestCategory) (err error)
	DeleteCategory(categoryId string) (err error)
	IsExistCategory(categoryId string) bool
}
