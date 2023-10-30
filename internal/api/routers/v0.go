package routers

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/handlers"
	"Sesuai/internal/api/middlewares"
	"github.com/kataras/iris/v12"
)

func Init(app *constracts.App, crs iris.Handler) {

	v1 := app.Iris.Party("/v1", crs)
	{
		// auth
		v1.Post("/auth/check-email", handlers.CheckEmail)
		v1.Post("/auth/login", handlers.Login)
		v1.Post("/auth/register", handlers.Register)

		// user
		v1.Get("/user", middlewares.AuthAdmin, handlers.GetUser)

		// blood type
		v1.Get("/blood-type", handlers.BloodType)

		// generate token
		v1.Get("/generate-token", middlewares.AuthAdmin, handlers.GetGenerateToken)
		v1.Post("/generate-token", middlewares.AuthAdmin, handlers.GenerateNewToken)

		// element
		v1.Get("/element", middlewares.AuthAdmin, handlers.GetElements)
		v1.Get("/element/{elementId}", middlewares.AuthAdmin, handlers.GetElementDetail)
		v1.Post("/element", middlewares.AuthAdmin, handlers.SaveElement)
		v1.Put("/element/{elementId}", middlewares.AuthAdmin, handlers.UpdateElement)
		v1.Delete("/element/{elementId}", middlewares.AuthAdmin, handlers.DeleteElement)

		// question
		v1.Get("/question", middlewares.AuthAdmin, handlers.GetQuestions)
		v1.Get("/question/all/{elementId}", middlewares.AuthAdmin, handlers.GetAllQuestionsByElementId)
		v1.Get("/question/{questionId}", middlewares.AuthAdmin, handlers.GetQuestion)
		v1.Post("/question", middlewares.AuthAdmin, handlers.SaveQuestion)
		v1.Put("/question/{questionId}", middlewares.AuthAdmin, handlers.UpdateQuestion)
		v1.Delete("/question/{questionId}", middlewares.AuthAdmin, handlers.DeleteQuestion)

		// submission
		v1.Get("/submissions", middlewares.AuthAdmin, handlers.GetSubmissions)
		v1.Get("/submissions/result/{submissionId}", middlewares.AuthAdmin, handlers.GetResultSubmission)

		// horoscope point
		v1.Get("/horoscope-point/{elementId}", middlewares.AuthAdmin, handlers.GetHoroscopePoint)
		v1.Put("/horoscope-point", middlewares.AuthAdmin, handlers.UpdateHoroscopePoint)

		// shio point
		v1.Get("/shio-point/{elementId}", middlewares.AuthAdmin, handlers.GetShioPoint)
		v1.Put("/shio-point", middlewares.AuthAdmin, handlers.UpdateShioPoint)

		// blood type point
		v1.Get("/blood-type-point/{elementId}", middlewares.AuthAdmin, handlers.GetBloodTypePoint)
		v1.Put("/blood-type-point", middlewares.AuthAdmin, handlers.UpdateBloodTypePoint)

		// point answer
		v1.Get("/point-answer", middlewares.AuthAdmin, handlers.GetPointAnswer)
		v1.Put("/point-answer", middlewares.AuthAdmin, handlers.UpdatePointAnswer)

		// admin list
		v1.Get("/admins", middlewares.AuthAdmin, handlers.GetAdmins)
		v1.Get("/admins/{adminId}", middlewares.AuthAdmin, handlers.GetAdmin)
		v1.Post("/admins", middlewares.AuthAdmin, handlers.SaveAdmin)
		v1.Put("/admins/{adminId}", middlewares.AuthAdmin, handlers.UpdateAdmin)
		v1.Delete("/admins/{adminId}", middlewares.AuthAdmin, handlers.DeleteAdmin)

		// access menu by admin id
		v1.Get("/access-menu/{adminId}", middlewares.AuthAdmin, handlers.GetAccessMenu)

		// check user for question test
		v1.Get("/check-question-test", middlewares.AuthUser, handlers.CheckQuestionTest)

		// question test
		v1.Get("/question-test", middlewares.AuthUser, handlers.GetQuestionsTest)

		// submit question test
		v1.Post("/submit-question-test", middlewares.AuthUser, handlers.SubmitQuestionTest)

		// result
		v1.Get("/result", middlewares.AuthUser, handlers.GetResult)
		v1.Get("/result/all", middlewares.AuthUser, handlers.GetAllResult)

		// profile user
		v1.Get("/profile-user", middlewares.AuthUser, handlers.GetProfileUser)
	}
}
