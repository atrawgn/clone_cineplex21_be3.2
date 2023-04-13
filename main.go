package main

import (
	"github.com/atrawiguna/golang-restapi-gorm/database"
	"github.com/atrawiguna/golang-restapi-gorm/migration"
	"github.com/atrawiguna/golang-restapi-gorm/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()
	//CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// INITIAL ROUTE
	route.RouteInit(app)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
