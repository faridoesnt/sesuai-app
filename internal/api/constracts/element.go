package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type ElementRepository interface {
	FindElements() (elements []entities.Element, err error)
	FindElementById(elementId string) (element entities.Element, err error)
	InsertElement(element entities.RequestElement) (err error)
	UpdateElement(elementId string, params entities.RequestElement) (err error)
	DeleteElement(elementId string) (err error)
}

type ElementService interface {
	GetElements() (listElement []response.Element)
	GetElementDetail(elementId string) (element response.Element)
	InsertElement(element entities.RequestElement) (err error)
	UpdateElement(elementId string, params entities.RequestElement) (err error)
	DeleteElement(elementId string) (err error)
	IsExistElement(elementId string) bool
}
