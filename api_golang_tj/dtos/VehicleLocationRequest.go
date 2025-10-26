package dtos

type VehicleLocation struct {
	VehicleID string  `json:"vehicle_id" db:"vehicle_id"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Timestamp int64   `json:"timestamp" db:"timestamp"`
}
