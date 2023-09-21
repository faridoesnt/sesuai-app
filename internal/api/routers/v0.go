package routers

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/handlers"
	"Sesuai/internal/api/middlewares"
	"github.com/kataras/iris/v12"
)

func Init(app *constracts.App, crs iris.Handler) {

	app.Iris.Get("/healthcheck", func(c iris.Context) {
		_, _ = c.WriteString("OK")
	})

	v1 := app.Iris.Party("/v1", crs)
	{
		// auth
		v1.Post("/auth/check-email", handlers.CheckEmail)
		v1.Post("/auth/login", handlers.Login)
		v1.Post("/auth/register", handlers.Register)

		// user
		v1.Get("/user", middlewares.Auth, handlers.GetUser)

		// blood type
		v1.Get("/blood-type", handlers.BloodType)

		// generate token
		v1.Get("/generate-token", middlewares.Auth, handlers.GetGenerateToken)
		v1.Post("/generate-token", middlewares.Auth, handlers.GenerateNewToken)
		v1.Post("/generate-token/use", middlewares.Auth, handlers.UseToken)

		// category
		v1.Get("/category", middlewares.Auth, handlers.GetCategory)
		v1.Get("/category/{categoryId}", middlewares.Auth, handlers.GetCategoryDetail)
		v1.Post("/category", middlewares.Auth, handlers.SaveCategory)
		v1.Put("/category/{categoryId}", middlewares.Auth, handlers.UpdateCategory)
		v1.Delete("/category/{categoryId}", middlewares.Auth, handlers.DeleteCategory)

		// question
		v1.Get("/question", middlewares.Auth, handlers.GetQuestions)
		v1.Get("/question/{questionId}", middlewares.Auth, handlers.GetQuestion)
		v1.Post("/question", middlewares.Auth, handlers.SaveQuestion)
		v1.Delete("/question/{questionId}", middlewares.Auth, handlers.DeleteQuestion)

		// submission
		v1.Get("/submissions", middlewares.Auth, handlers.GetSubmissions)
		v1.Get("/submissions/result/{submissionId}", middlewares.Auth, handlers.GetResultSubmission)
	}
}
