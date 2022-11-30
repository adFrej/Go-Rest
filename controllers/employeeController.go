package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goRest/databse"
	"goRest/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getEmployeesFilter(filter bson.D, options *options.FindOptions) (employees []models.Employee, err error) {
	coll := databse.Client.Database("company").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	cursor, err := coll.Find(ctx, filter, options)
	cancel()
	if err != nil {
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 300*time.Second)
	err = cursor.All(ctx, &employees)
	cancel()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func GetEmployees(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Query().Has("city") {
		return GetEmployeesByCity(w, r)
	}
	log.Println("Request for all employees.")
	defer log.Println("Request completed.")
	employees, err := getEmployeesFilter(bson.D{}, options.Find().SetLimit(100))
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employees)
	return err
}

func GetEmployeesByCity(w http.ResponseWriter, r *http.Request) error {
	var city = r.URL.Query().Get("city")
	if city == "" {
		city = mux.Vars(r)["city"]
	}
	log.Println("Request for employees with city:", city)
	defer log.Println("Request completed.")
	employees, err := getEmployeesFilter(bson.D{{"city", bson.D{{"$eq", city}}}}, options.Find())
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employees)
	return err
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) error {
	log.Println("Request for employee with id:", mux.Vars(r)["id"])
	defer log.Println("Request completed.")
	employeeId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode("Bad id format.")
			return err
		}
		return err
	}
	coll := databse.Client.Database("company").Collection("employees")
	var employee models.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	err = coll.FindOne(ctx, bson.D{{"id", employeeId}}).Decode(&employee)
	cancel()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode("Employee not found.")
			return err
		}
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employee)
	return err
}
