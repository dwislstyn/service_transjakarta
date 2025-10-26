package rabbit

import (
	"encoding/json"
	"log"
	"os"

	"example.com/service_transjakarta/api_golang_tj/dtos"
	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel

func ConnectRabbit() {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		log.Println("❌ Gagal konek RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("❌ Gagal buka channel RabbitMQ:", err)
	}

	// Declare exchange
	err = ch.ExchangeDeclare(
		"fleet.events",
		"fanout",
		true, false, false, false, nil,
	)
	if err != nil {
		log.Println("❌ Gagal declare exchange:", err)
	}

	// Declare queue utamanya
	queue, err := ch.QueueDeclare(
		"geofence_alerts", // nama queue
		true,              // durable
		false,             // autoDelete
		false,             // exclusive
		false,             // noWait
		nil,
	)
	if err != nil {
		log.Println("❌ Gagal declare queue:", err)
	}

	// Bind queue ke exchange fleet.events
	err = ch.QueueBind(
		queue.Name,     // queue
		"",             // routing key (kosong karena fanout)
		"fleet.events", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Println("❌ Gagal bind queue:", err)
	}

	RabbitChannel = ch
	log.Println("✅ RabbitMQ Ready!")
}

func PublishGeofenceEvent(loc dtos.VehicleLocation) {
	if RabbitChannel == nil {
		log.Println("⚠ RabbitMQ belum siap")
	}

	// Format payload sesuai requirement
	payload := map[string]interface{}{
		"vehicle_id": loc.VehicleID,
		"event":      "geofence_entry",
		"location": map[string]float64{
			"latitude":  loc.Latitude,
			"longitude": loc.Longitude,
		},
		"timestamp": loc.Timestamp,
	}

	body, _ := json.Marshal(payload)

	err := RabbitChannel.Publish(
		"fleet.events",
		"", // routing key kosong karena tipe fanout
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("❌ Gagal publish ke RabbitMQ:", err)
	}

	log.Printf("📡 Event geofence_entry dikirim untuk kendaraan %s\n", loc.VehicleID)
}
