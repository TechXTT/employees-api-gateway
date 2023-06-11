package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/data/pb"
	"github.com/gorilla/mux"
)

type CreatePositionRequestBody struct {
	Title        string `json:"title"`
	DepartmentID string `json:"department_id"`
}

type UpdatePositionRequestBody struct {
	Title        string `json:"title"`
	DepartmentID string `json:"department_id"`
}

func GetPositions(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	res, err := c.GetPositions(context.Background(), &pb.GetPositionsRequest{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetPosition(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	id := mux.Vars(r)["id"]

	res, err := c.GetPosition(context.Background(), &pb.GetPositionRequest{
		Id: id,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CreatePosition(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	var body CreatePositionRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := c.CreatePosition(context.Background(), &pb.CreatePositionRequest{
		Title:        body.Title,
		DepartmentId: body.DepartmentID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func UpdatePosition(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	var body UpdatePositionRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	res, err := c.UpdatePosition(context.Background(), &pb.UpdatePositionRequest{
		Id:           id,
		Title:        body.Title,
		DepartmentId: body.DepartmentID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeletePosition(w http.ResponseWriter, r *http.Request, c pb.DataServiceClient) {
	id := mux.Vars(r)["id"]

	_, err := c.DeletePosition(context.Background(), &pb.DeletePositionRequest{
		Id: id,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
