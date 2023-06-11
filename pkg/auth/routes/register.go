package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	var body RegisterRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[Register] Error: ", err)
		return
	}

	if res.Status != http.StatusOK {
		w.WriteHeader(int(res.Status))
		return
	}

	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
