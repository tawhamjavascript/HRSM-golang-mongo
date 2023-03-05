package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhamjavascript/HRSM/routers"
	"log"
)

func main() {
	app := fiber.New()
	routers.RoutesEmployees(app)
	log.Fatal(app.Listen(":8000"))

}
