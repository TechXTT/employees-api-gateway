package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/data/pb"
	"github.com/gorilla/mux"
)

type CreateDepartmentRequestBody struct {
	Name string `json:"name"`
}

type UpdateDepartmentRequestBody struct {
	Name string `json:"name"`
}

func GetDepartments(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	res, err := c.GetDepartments(context.Background(), &pb.GetDepartmentsRequest{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetDepartment(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	id := mux.Vars(r)["id"]

	res, err := c.GetDepartment(context.Background(), &pb.GetDepartmentRequest{
		Id: id,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreateDepartment(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	var body CreateDepartmentRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := c.CreateDepartment(context.Background(), &pb.CreateDepartmentRequest{
		Name: body.Name,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdateDepartment(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	var body UpdateDepartmentRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	res, err := c.UpdateDepartment(context.Background(), &pb.UpdateDepartmentRequest{
		Id:   id,
		Name: body.Name,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	id := mux.Vars(r)["id"]

	_, err := c.DeleteDepartment(context.Background(), &pb.DeleteDepartmentRequest{
		Id: id,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
