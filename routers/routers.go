package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhamjavascript/HRSM/controllers"
)

func RoutesEmployees(app *fiber.App) *fiber.App {
	app.Get("/employees", controllers.GetAll)
	app.Post("/employees", controllers.Add)
	app.Put("/employees/:id", controllers.Update)
	app.Delete("/employees/:id", controllers.Delete)
	return app
}
