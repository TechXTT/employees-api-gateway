package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/data/pb"
	"github.com/gorilla/mux"
)

type CreateEmployeeRequestBody struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Email      string  `json:"email"`
	HireDate   string  `json:"hireDate"`
	Salary     float32 `json:"salary"`
	PositionId string  `json:"positionId"`
}

type UpdateEmployeeRequestBody struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Email      string  `json:"email"`
	HireDate   string  `json:"hireDate"`
	Salary     float32 `json:"salary"`
	PositionId string  `json:"positionId"`
}

func GetEmployees(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	res, err := c.GetEmployees(context.Background(), &pb.GetEmployeesRequest{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetEmployee(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	id := mux.Vars(r)["id"]

	res, err := c.GetEmployee(context.Background(), &pb.GetEmployeeRequest{
		Id: id,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	var body CreateEmployeeRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := c.CreateEmployee(context.Background(), &pb.CreateEmployeeRequest{
		FirstName:  body.FirstName,
		LastName:   body.LastName,
		Email:      body.Email,
		HireDate:   body.HireDate,
		Salary:     body.Salary,
		PositionId: body.PositionId,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	var body UpdateEmployeeRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	res, err := c.UpdateEmployee(context.Background(), &pb.UpdateEmployeeRequest{
		Id:         id,
		FirstName:  body.FirstName,
		LastName:   body.LastName,
		Email:      body.Email,
		HireDate:   body.HireDate,
		Salary:     body.Salary,
		PositionId: body.PositionId,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	id := mux.Vars(r)["id"]

	_, err := c.DeleteEmployee(context.Background(), &pb.DeleteEmployeeRequest{
		Id: id,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
