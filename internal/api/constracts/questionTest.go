package constracts

import "Sesuai/internal/api/entities"

type QuestionTestRepository interface {
	FindQuestionsTest() (questionsTest []entities.QuestionTest, err error)
}

type QuestionTestService interface {
	GetQuestionsTest() (questionsTest []entities.QuestionTest, err error)
}
