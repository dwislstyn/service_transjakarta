package controller

import (
	"net/http"
	"strconv"

	"example.com/service_transjakarta/api_golang_tj/exceptions"
	"example.com/service_transjakarta/api_golang_tj/usecase"
	"github.com/gorilla/mux"
)

type LocationController struct {
	LocationUseCase usecase.LocationUseCase
}

type ILocationController interface {
	InquiryLocation(response http.ResponseWriter, request *http.Request)
	InquiryListLocation(response http.ResponseWriter, request *http.Request)
}

func (l *LocationController) InquiryLocation(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	vehicleID := vars["vehicle_id"]

	location, err := l.LocationUseCase.InquiryLocation(vehicleID)
	if err != nil {
		exceptions.InvalidException(response, "Error InquiryLocation Use Case", err)
		return
	}

	if location == nil {
		exceptions.DataNotFoundException(response, "Data kendaraan dengan ID "+vehicleID+" tidak ditemukan", err)
		return
	}

	exceptions.SuccessResponse(response, "Inquiry current location success", location)
}

func (l *LocationController) InquiryListLocation(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	vehicleID := vars["vehicle_id"]

	startStr := request.URL.Query().Get("start")
	endStr := request.URL.Query().Get("end")

	// Parse UNIX timestamp (detik)
	startInt, _ := strconv.ParseInt(startStr, 10, 64)
	endInt, _ := strconv.ParseInt(endStr, 10, 64)

	location, err := l.LocationUseCase.InquiryListLocation(vehicleID, startInt, endInt)
	if err != nil {
		exceptions.InvalidException(response, "Error InquiryListLocation Use Case", err)
		return
	}

	if len(location) == 0 {
		exceptions.DataNotFoundException(response, "Data kendaraan dengan ID "+vehicleID+" tidak ditemukan", err)
		return
	}

	exceptions.SuccessResponse(response, "Inquiry list history location success", location)
}

func NewLocationController(usecase *usecase.LocationUseCase) LocationController {
	return LocationController{
		LocationUseCase: *usecase,
	}
}
