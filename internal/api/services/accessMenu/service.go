package accessMenu

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/constracts"
)

type Service struct {
	app  constracts.App
	repo constracts.AccessMenuRepository
}

func Init(a *constracts.App) (svc constracts.AccessMenuService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetAccessMenuByAdminId(adminId string) (accessMenus []string, err error) {
	menus, err := s.repo.FindAccessMenuByAdminId(adminId)

	if len(menus) > 0 {
		for _, menu := range menus {
			if menu.MenuName != "" {
				var result string

				switch menu.MenuName {
				case constants.GenerateToken:
					result = constants.EnumGenerateToken
				case constants.Element:
					result = constants.EnumElement
				case constants.QuestionList:
					result = constants.EnumQuestionList
				case constants.Submition:
					result = constants.EnumSubmition
				case constants.PointAnswer:
					result = constants.EnumPointAnswer
				case constants.AdminList:
					result = constants.EnumAdminList
				}
				accessMenus = append(accessMenus, result)
			}
		}
	}

	return
}
