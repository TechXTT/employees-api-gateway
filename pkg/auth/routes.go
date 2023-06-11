package auth

import (
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/auth/routes"
	"github.com/TechXTT/employees-api-gateway/pkg/config"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/login", svc.Login).Methods("POST")
	router.HandleFunc("/register", svc.Register).Methods("POST")

	return svc
}

func (svc *ServiceClient) Register(w http.ResponseWriter, r *http.Request) {
	routes.Register(w, r, svc.Client)
}

func (svc *ServiceClient) Login(w http.ResponseWriter, r *http.Request) {
	routes.Login(w, r, svc.Client)
}
