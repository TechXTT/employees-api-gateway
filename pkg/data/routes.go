package data

import (
	"net/http"

	"github.com/TechXTT/employees-api-gateway/pkg/auth"
	"github.com/TechXTT/employees-api-gateway/pkg/config"
	"github.com/TechXTT/employees-api-gateway/pkg/data/routes"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	router := r.PathPrefix("/data").Subrouter()
	router.Use(a.AuthRequired)
	router.HandleFunc("/employees/{id}", svc.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", svc.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", svc.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", svc.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", svc.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/positions/{id}", svc.GetPosition).Methods("GET")
	router.HandleFunc("/positions", svc.GetPositions).Methods("GET")
	router.HandleFunc("/positions", svc.CreatePosition).Methods("POST")
	router.HandleFunc("/positions/{id}", svc.UpdatePosition).Methods("PUT")
	router.HandleFunc("/positions/{id}", svc.DeletePosition).Methods("DELETE")
	router.HandleFunc("/departments/{id}", svc.GetDepartment).Methods("GET")
	router.HandleFunc("/departments", svc.GetDepartments).Methods("GET")
	router.HandleFunc("/departments", svc.CreateDepartment).Methods("POST")
	router.HandleFunc("/departments/{id}", svc.UpdateDepartment).Methods("PUT")
	router.HandleFunc("/departments/{id}", svc.DeleteDepartment).Methods("DELETE")
}

func (svc *ServiceClient) GetEmployees(w http.ResponseWriter, r *http.Request) {
	routes.GetEmployees(w, r, svc.Client)
}

func (svc *ServiceClient) GetEmployee(w http.ResponseWriter, r *http.Request) {
	routes.GetEmployee(w, r, svc.Client)
}

func (svc *ServiceClient) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	routes.CreateEmployee(w, r, svc.Client)
}

func (svc *ServiceClient) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	routes.UpdateEmployee(w, r, svc.Client)
}

func (svc *ServiceClient) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	routes.DeleteEmployee(w, r, svc.Client)
}

func (svc *ServiceClient) GetPositions(w http.ResponseWriter, r *http.Request) {
	routes.GetPositions(w, r, svc.Client)
}

func (svc *ServiceClient) GetPosition(w http.ResponseWriter, r *http.Request) {
	routes.GetPosition(w, r, svc.Client)
}

func (svc *ServiceClient) CreatePosition(w http.ResponseWriter, r *http.Request) {
	routes.CreatePosition(w, r, svc.Client)
}

func (svc *ServiceClient) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	routes.UpdatePosition(w, r, svc.Client)
}

func (svc *ServiceClient) DeletePosition(w http.ResponseWriter, r *http.Request) {
	routes.DeletePosition(w, r, svc.Client)
}

func (svc *ServiceClient) GetDepartments(w http.ResponseWriter, r *http.Request) {
	routes.GetDepartments(w, r, svc.Client)
}

func (svc *ServiceClient) GetDepartment(w http.ResponseWriter, r *http.Request) {
	routes.GetDepartment(w, r, svc.Client)
}

func (svc *ServiceClient) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	routes.CreateDepartment(w, r, svc.Client)
}

func (svc *ServiceClient) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	routes.UpdateDepartment(w, r, svc.Client)
}

func (svc *ServiceClient) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	routes.DeleteDepartment(w, r, svc.Client)
}
