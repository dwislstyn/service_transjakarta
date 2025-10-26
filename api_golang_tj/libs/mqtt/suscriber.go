package mqtt

import (
	"encoding/json"
	"log"

	"example.com/service_transjakarta/api_golang_tj/database/rabbit"
	"example.com/service_transjakarta/api_golang_tj/dtos"
	"example.com/service_transjakarta/api_golang_tj/libs/geofrance"
	"example.com/service_transjakarta/api_golang_tj/repositories"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(client mqtt.Client, repo *repositories.LocationRepository) {
	topic := "/fleet/vehicle/+/location"

	token := client.Subscribe(topic, 1, messageHandler(repo))
	token.Wait()

	if token.Error() != nil {
		log.Println("❌ MQTT Subscribe Failed:", token.Error())
		return
	}

	log.Println("✅ MQTT Subscriber aktif pada topic:", topic)
}

func messageHandler(repo *repositories.LocationRepository) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		log.Println("📩 Pesan diterima pada topic:", msg.Topic())

		var location dtos.VehicleLocation
		if err := json.Unmarshal(msg.Payload(), &location); err != nil {
			log.Println("❌ Error parse JSON:", err)
			return
		}

		log.Printf("🚗 Kendaraan %s Lokasi: (%.6f, %.6f) - %d\n",
			location.VehicleID, location.Latitude, location.Longitude, location.Timestamp)

		// Simpan ke DB
		if err := repo.InsertLocation(location); err != nil {
			log.Println("❌ Gagal menyimpan ke DB:", err)
		}

		dist := geofrance.CalculateDistance(location.Latitude, location.Longitude, geofrance.GeofenceLat, geofrance.GeofenceLon)

		if dist <= geofrance.GeofenceRadiusM {
			if !geofrance.VehicleState[location.VehicleID] {
				geofrance.VehicleState[location.VehicleID] = true
				rabbit.PublishGeofenceEvent(location)
			}
		} else {
			geofrance.VehicleState[location.VehicleID] = false
		}
	}
}
