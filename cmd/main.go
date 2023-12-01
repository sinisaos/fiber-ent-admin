package main

import (
	"log"

	"github.com/sinisaos/fiber-ent-admin/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "X-Total-Count",
	}))

	app.Static("/", "./admin/dist", fiber.Static{
		Index: "index.html",
	})

	router.SetupRoutes(app)

	app.Static("/swagger", "./api/docs", fiber.Static{
		Index: "index.html",
	})

	log.Fatal(app.Listen(":3000"))
}
