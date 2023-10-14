package element

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.ElementRepository
}

func Init(a *constracts.App) (svc constracts.ElementService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetElements() (listElement []response.Element) {
	elements, _ := s.repo.FindElements()

	for _, element := range elements {
		listElement = append(listElement, response.Element{
			Id:    element.Id,
			Name:  element.Name,
			Photo: element.Photo,
		})
	}

	return
}

func (s Service) InsertElement(element entities.RequestElement) (err error) {
	err = s.repo.InsertElement(element)

	return
}

func (s Service) GetElementDetail(elementId string) (element response.Element) {
	data, _ := s.repo.FindElementById(elementId)

	element.Id = data.Id
	element.Name = data.Name
	element.Photo = data.Photo

	return
}

func (s Service) UpdateElement(elementId string, params entities.RequestElement) (err error) {
	err = s.repo.UpdateElement(elementId, params)

	return
}

func (s Service) DeleteElement(elementId string) (err error) {
	err = s.repo.DeleteElement(elementId)

	return
}

func (s Service) IsExistElement(elementId string) bool {
	_, err := s.repo.FindElementById(elementId)
	if err != nil {
		return false
	}

	return true
}
