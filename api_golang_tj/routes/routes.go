package routes

import (
	"example.com/service_transjakarta/api_golang_tj/controller"
	"example.com/service_transjakarta/api_golang_tj/repositories"
	"example.com/service_transjakarta/api_golang_tj/usecase"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(db *sqlx.DB) *mux.Router {
	locationRepo := repositories.NewLocationRepository(db)
	locationUseCase := usecase.NewLocationUseCase(&locationRepo)
	locationController := controller.NewLocationController(&locationUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/vehicles/{vehicle_id}/location", locationController.InquiryLocation).Methods("GET")
	r.HandleFunc("/vehicles/{vehicle_id}/history", locationController.InquiryListLocation).Methods("GET")

	return r
}
