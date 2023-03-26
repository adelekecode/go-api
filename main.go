package main

import (
	"log"

	"github.com/adeleke-code/go_api/database"
	"github.com/adeleke-code/go_api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my first API with GO")
}


func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("", welcome)
	// User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.GET("/api/users/:id", routes.Getuser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	
}
func main() {
	database.ConnectDb()
    app := fiber.New()

	setupRoutes(app)


    log.Fatal(app.Listen(":3000"))
}