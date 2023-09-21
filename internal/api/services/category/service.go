package category

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.CategoryRepository
}

func Init(a *constracts.App) (svc constracts.CategoryService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetCategory() (listCategory []response.Category) {
	categories, _ := s.repo.FindCategory()

	for _, category := range categories {
		listCategory = append(listCategory, response.Category{
			Id:    category.Id,
			Name:  category.Name,
			Photo: category.Photo,
		})
	}

	return
}

func (s Service) InsertCategory(category entities.RequestCategory) (err error) {
	err = s.repo.InsertCategory(category)

	return
}

func (s Service) GetCategoryDetail(categoryId string) (category response.Category) {
	data, _ := s.repo.FindCategoryById(categoryId)

	category.Id = data.Id
	category.Name = data.Name
	category.Photo = data.Photo

	return
}

func (s Service) UpdateCategory(categoryId string, params entities.RequestCategory) (err error) {
	err = s.repo.UpdateCategory(categoryId, params)

	return
}

func (s Service) DeleteCategory(categoryId string) (err error) {
	err = s.repo.DeleteCategory(categoryId)

	return
}

func (s Service) IsExistCategory(categoryId string) bool {
	_, err := s.repo.FindCategoryById(categoryId)
	if err != nil {
		return false
	}

	return true
}
