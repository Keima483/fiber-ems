package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/fiber-ems/controllers"
	"github.com/keima483/fiber-ems/repository"
)

func main() {
	if err := repository.InitialiseDB(); err != nil {
		panic(err.Error())
	}
	app := fiber.New()
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
	employeeAPI.Post("/:id", controllers.AddEmployee)
	employeeAPI.Get("/:id", controllers.GetEmployees)
	employeeAPI.Put("", controllers.UpdateEmployee)
	employeeAPI.Delete("/:id", controllers.DeleteEmployee)
}