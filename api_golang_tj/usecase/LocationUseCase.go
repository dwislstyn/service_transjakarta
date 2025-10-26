package usecase

import (
	"example.com/service_transjakarta/api_golang_tj/dtos"
	"example.com/service_transjakarta/api_golang_tj/repositories"
)

type LocationUseCase struct {
	LocationRepo repositories.LocationRepository
}

type ILocationUseCase interface {
	InquiryLocation(data string) (*dtos.VehicleLocation, error)
	InquiryListLocation(vehicleId string, start int64, end int64) ([]dtos.VehicleLocation, error)
}

func (l *LocationUseCase) InquiryLocation(vehicleId string) (*dtos.VehicleLocation, error) {
	var results *dtos.VehicleLocation
	getCurrentLocation, err := l.LocationRepo.GetCurrentLocation(vehicleId)

	if err != nil || getCurrentLocation == nil {
		return results, err
	}

	return getCurrentLocation, err
}

func (l *LocationUseCase) InquiryListLocation(vehicleId string, start int64, end int64) ([]dtos.VehicleLocation, error) {
	var results []dtos.VehicleLocation
	getListLocation, err := l.LocationRepo.GetListLocation(vehicleId, start, end)

	if err != nil {
		return results, err
	}

	defer getListLocation.Close()

	for getListLocation.Next() {
		var location dtos.VehicleLocation
		if err := getListLocation.StructScan(&location); err != nil {
			return results, err
		}

		results = append(results, location)
	}

	if err := getListLocation.Err(); err != nil {
		return results, err
	}

	return results, err
}

func NewLocationUseCase(repo *repositories.LocationRepository) LocationUseCase {
	return LocationUseCase{
		LocationRepo: *repo,
	}
}
