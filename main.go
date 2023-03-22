package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/fiber-ems/controllers"
	"github.com/keima483/fiber-ems/initializers"
	"github.com/keima483/fiber-ems/middleware"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic(err.Error())
	}
	if err := initializers.InitialiseDB(&config); err != nil {
		panic(err.Error())
	}
}

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("refer to https://github.com/Keima483/fiber-ems for the endpoints")
	})
	companyRoutes(app)
	employeeRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		panic("Post :3000 already taken")
	}
}

func companyRoutes(app *fiber.App) {
	companyAPI := app.Group("/api/v1/company")
	companyAPI.Post("/login", controllers.Login)
	companyAPI.Post("/signup", controllers.Signup)
}

func employeeRoutes(app *fiber.App) {
	employeeAPI := app.Group("/api/v1/employee")
	employeeAPI.Use(middleware.VerifyJWTToken)

	employeeAPI.Post("", controllers.AddEmployee)
	employeeAPI.Get("", controllers.GetEmployees)
	employeeAPI.Put("", controllers.UpdateEmployee)
	employeeAPI.Delete("/:id", controllers.DeleteEmployee)
}
