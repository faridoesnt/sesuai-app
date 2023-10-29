package constracts

import "Sesuai/internal/api/entities"

type QuestionTestRepository interface {
	FindQuestionsTest() (questionsTest []entities.QuestionTest, err error)
	SubmitQuestionTest(params entities.SubmitQuestionTest, userId string, totalPointQuestionsByElement map[string]float64) (err error)
	CountQuestionTestUser(userId string) (total int64, err error)
}

type QuestionTestService interface {
	GetQuestionsTest() (questionsTest []entities.QuestionTest, err error)
	SubmitQuestionTest(params entities.SubmitQuestionTest, userId string, totalPointQuestionsByElement map[string]float64) (err error)
	CheckQuestionTestUser(userId string) (isExist bool, err error)
}
