package geofrance

import "math"

const (
	GeofenceLat     = -6.2088 // Koordinat titik geofence
	GeofenceLon     = 106.8456
	GeofenceRadiusM = 50.0 // Radius 50 meter
)

// State kendaraan: true = sudah berada dalam radius
var VehicleState = make(map[string]bool)

// Haversine Formula untuk hitung jarak dalam meter
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3 // Earth radius in meter
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)

	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}
