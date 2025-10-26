package repositories

import (
	"database/sql"
	"fmt"

	"example.com/service_transjakarta/api_golang_tj/dtos"
	"github.com/jmoiron/sqlx"
)

type LocationRepository struct {
	*sqlx.DB
}

type ILocationRepository interface {
	GetListLocation(vehicleId string, start int64, end int64) (*sqlx.Rows, error)
	GetCurrentLocation(vehicleId string) (*dtos.VehicleLocation, error)
	InsertLocation(data dtos.VehicleLocation) error
}

var LocationResponse dtos.VehicleLocation

func (l *LocationRepository) GetListLocation(vehicleId string, start int64, end int64) (*sqlx.Rows, error) {
	query := `
		SELECT 
		vehicle_id,
		latitude,
		longitude,
		timestamp
		FROM vehicle_locations
		WHERE 
		vehicle_id = $1
		AND timestamp BETWEEN $2 AND $3
		ORDER BY timestamp DESC
	`

	rows, err := l.DB.Queryx(query, vehicleId, start, end)
	if err != nil {
		return nil, fmt.Errorf("gagal menjalankan query: %v", err)
	}

	return rows, nil
}

func (l *LocationRepository) GetCurrentLocation(vehicleId string) (*dtos.VehicleLocation, error) {
	row := l.DB.QueryRowx(`
		SELECT vehicle_id, latitude, longitude, timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`, vehicleId)

	location := &dtos.VehicleLocation{}
	err := row.StructScan(location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return location, nil
}

func (l *LocationRepository) InsertLocation(data dtos.VehicleLocation) error {
	query := `INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp)
			  VALUES (:vehicle_id, :latitude, :longitude, :timestamp)`

	_, err := l.DB.NamedExec(query, data)
	return err
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return LocationRepository{
		db,
	}
}
