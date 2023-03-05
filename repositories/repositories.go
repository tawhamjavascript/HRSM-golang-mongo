package repositories

import (
	"fmt"
	"github.com/tawhamjavascript/HRSM/db"
	"github.com/tawhamjavascript/HRSM/models"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllEmployees(context *fasthttp.RequestCtx) ([]models.Employee, error) {
	mongo := db.Connect()
	query := bson.D{{}}
	employees := make([]models.Employee, 0)
	cursor, err := mongo.Db.Collection("employees").Find(context, &query)
	if err != nil {
		return employees, err
	}
	err = cursor.All(context, &employees)
	if err != nil {
		return employees, err

	}
	return employees, err
}

func AddEmployee(context *fasthttp.RequestCtx, employee *models.Employee) (*models.Employee, error) {

	mg := db.Connect()
	collection := mg.Db.Collection("employees")
	employee.ID = ""
	insertionResult, err := collection.InsertOne(context, employee)
	if err != nil {
		return employee, err
	}
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(context, filter)
	fmt.Println(createdRecord)
	createEmployee := &models.Employee{}
	createdRecord.Decode(&createEmployee)

	return createEmployee, nil
}

func UpdateEmployee(context *fasthttp.RequestCtx, idParam string, employee *models.Employee) (*models.Employee, error) {
	EmployeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return employee, err
	}

	query := bson.D{{Key: "_id", Value: EmployeeID}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}
	mg := db.Connect()
	err = mg.Db.Collection("employees").FindOneAndUpdate(context, &query, &update).Err()
	if err != nil {
		return employee, err
	}
	employee.ID = idParam
	return employee, nil

}

func DeleteEmployee(context *fasthttp.RequestCtx, idParam string) (string, error) {
	mg := db.Connect()
	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return "", err
	}
	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := mg.Db.Collection("employees").DeleteOne(context, &query)
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	if result.DeletedCount < 1 {
		return "Failed", nil

	}
	return "ok", nil
}
