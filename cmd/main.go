package main

import (
	"log"
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/auth"
	"github.com/TechXTT/employees-api-gateway/pkg/config"
	"github.com/TechXTT/employees-api-gateway/pkg/data"
	"github.com/gorilla/mux"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := mux.NewRouter()

	authSvc := *auth.RegisterRoutes(r, c)
	data.RegisterRoutes(r, c, &authSvc)

	log.Fatal(http.ListenAndServe(c.Port, r))
}
