package router

import (
	"github.com/sinisaos/fiber-ent-admin/pkg/database"
	"github.com/sinisaos/fiber-ent-admin/pkg/handler"
	"github.com/sinisaos/fiber-ent-admin/pkg/middleware"
	"github.com/sinisaos/fiber-ent-admin/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Setup api routes
func SetupRoutes(app *fiber.App) {
	// DB client
	client := database.DbConnection()

	// Services
	questionService := service.NewQuestionService(client)
	answerService := service.NewAnswerService(client)
	userService := service.NewUserService(client)
	authService := service.NewAuthService(client)
	tagService := service.NewTagService(client)

	// Handlers
	questionHandler := handler.NewQuestionHandler(*questionService)
	answerHandler := handler.NewAnswerHandler(*answerService)
	userHandler := handler.NewUserHandler(*userService)
	authHandler := handler.NewAuthHandler(*authService)
	tagHandler := handler.NewTagHandler(*tagService)

	// Logger middleware
	api := app.Group("/", logger.New())

	// Auth route
	api.Post("/login", authHandler.LoginHandler)

	// Users routes
	user := api.Group("/users")
	user.Get("/", userHandler.GetAllUsersHandler)
	user.Get("/:id", userHandler.GetUserHandler)
	user.Get("/:id/answers", userHandler.GetUserAnswersHandler)
	user.Get("/:id/questions", userHandler.GetUserQuestionsHandler)
	user.Post("/", middleware.AuthMiddleware(), userHandler.CreateUserHandler)
	user.Put("/:id", middleware.AuthMiddleware(), userHandler.UpdateUserHandler)
	user.Delete("/:id", middleware.AuthMiddleware(), userHandler.DeleteUserHandler)

	// Questions routes
	question := api.Group("/questions")
	question.Get("/", questionHandler.GetAllQuestionsHandler)
	question.Get("/:id", questionHandler.GetQuestionHandler)
	question.Get("/:id/answers", questionHandler.GetQuestionAnswersHandler)
	question.Get("/:id/author", questionHandler.GetQuestionAuthorHandler)
	question.Get("/:id/tags", questionHandler.GetQuestionTagsHandler)
	question.Post("/", middleware.AuthMiddleware(), questionHandler.CreateQuestionHandler)
	question.Put("/:id", middleware.AuthMiddleware(), questionHandler.UpdateQuestionHandler)
	question.Delete("/:id", middleware.AuthMiddleware(), questionHandler.DeleteQuestionHandler)

	// Answer
	answer := api.Group("/answers")
	answer.Get("/", answerHandler.GetAllAnswersHandler)
	answer.Get("/:id", answerHandler.GetAnswerHandler)
	answer.Get("/:id/author", answerHandler.GetAnswerAuthorHandler)
	answer.Get("/:id/question", answerHandler.GetAnswerQuestionHandler)
	answer.Post("/", middleware.AuthMiddleware(), answerHandler.CreateAnswerHandler)
	answer.Put("/:id", middleware.AuthMiddleware(), answerHandler.UpdateAnswerHandler)
	answer.Delete("/:id", middleware.AuthMiddleware(), answerHandler.DeleteAnswerHandler)

	// Tag
	tag := api.Group("/tags")
	tag.Get("/", tagHandler.GetAllTagsHandler)
	tag.Get("/:id", tagHandler.GetTagHandler)
	tag.Get("/:id/question", tagHandler.GetTagQuestionHandler)
	tag.Post("/", middleware.AuthMiddleware(), tagHandler.CreateTagHandler)
	tag.Put("/:id", middleware.AuthMiddleware(), tagHandler.UpdateTagHandler)
	tag.Delete("/:id", middleware.AuthMiddleware(), tagHandler.DeleteTagHandler)
}
