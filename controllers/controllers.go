package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tawhamjavascript/HRSM/models"
	"github.com/tawhamjavascript/HRSM/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(c *fiber.Ctx) error {
	employees, err := repositories.GetAllEmployees(c.Context())

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)

}

func Add(c *fiber.Ctx) error {
	employee := &models.Employee{}

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	newEmployee, err := repositories.AddEmployee(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(newEmployee)
}

func Update(c *fiber.Ctx) error {
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	idParam := c.Params("id")
	updatedEmployee, err := repositories.UpdateEmployee(c.Context(), idParam, employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(updatedEmployee)
}

func Delete(c *fiber.Ctx) error {
	result, err := repositories.DeleteEmployee(c.Context(), c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}
	if result == "failed" {
		return c.SendStatus(404)
	}
	return c.Status(200).SendString("record deleted")

}
