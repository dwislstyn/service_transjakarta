package database

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"example.com/service_transjakarta/api_golang_tj/libs/geofrance"
	mqttLocal "example.com/service_transjakarta/api_golang_tj/libs/mqtt"
	"example.com/service_transjakarta/api_golang_tj/repositories"
	mqttGit "github.com/eclipse/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectMqtt() mqttGit.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found, using system environment variables")
	}

	mqttURL := os.Getenv("MQTT_BROKER_URL")
	opts := mqttGit.NewClientOptions().AddBroker(mqttURL)
	opts.SetClientID("vehicle-location-client")

	client := mqttGit.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("❌ MQTT Connection Failed:", token.Error())
	}

	log.Println("✅ Terhubung ke MQTT:", mqttURL)
	return *&client
}

func SubscribeMqtt(client mqttGit.Client, db *sqlx.DB) {
	repo := repositories.NewLocationRepository(db)
	mqttLocal.Subscribe(client, &repo)
}

func PublishMockLocation(client mqttGit.Client) {
	vehicleIDs := []string{"B1234XYZ", "B5678ABC", "C9012DEF", "D3456GHI", "E7890JKL"}

	rand.Seed(time.Now().UnixNano())

	offset := 0.00050
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			// Pilih vehicleID secara random
			vehicleID := vehicleIDs[rand.Intn(len(vehicleIDs))]

			latitude := geofrance.GeofenceLat + (rand.Float64()*offset - offset/2)
			longitude := geofrance.GeofenceLon + (rand.Float64()*offset - offset/2)

			payload := fmt.Sprintf(`{"vehicle_id":"%s","latitude":%.6f,"longitude":%.6f,"timestamp":%d}`,
				vehicleID, latitude, longitude, time.Now().Unix())

			token := client.Publish(fmt.Sprintf("/fleet/vehicle/%s/location", vehicleID), 0, false, payload)
			token.Wait()
			log.Println("Published:", payload)
		}
	}()
}

func Connect() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Default value kalau tidak diset
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, sslmode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("❌ gagal membuka koneksi ke DB: %v", err)
	}

	return db, nil
}
