package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	EmailOrUsername string `json:"emailorusername"`
	Password        string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	var body LoginRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		EmailOrUsername: body.EmailOrUsername,
		Password:        body.Password,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.Status != http.StatusOK {
		w.WriteHeader(int(res.Status))
		return
	}

	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
