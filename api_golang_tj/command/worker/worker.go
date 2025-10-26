package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	log.Println("🚀 Worker geofence START")

	// Siapkan file log
	fileLog, err := os.OpenFile("logs/worker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("❌ Gagal buka worker.log:", err)
	}
	defer fileLog.Close()
	log.SetOutput(fileLog)

	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found, using system environment variables")
	}

	// Koneksi ke RabbitMQ
	rabbitURL := os.Getenv("RABBITMQ_URL")
	fmt.Println(rabbitURL)
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		log.Fatal("❌ Gagal konek RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("❌ Gagal buka channel RabbitMQ:", err)
	}
	defer ch.Close()

	// Pastikan queue ter-declare (safety)
	queue, err := ch.QueueDeclare(
		"geofence_alerts",
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("❌ Gagal declare queue:", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("❌ Gagal consume:", err)
	}

	log.Println("✅ Worker siap menerima pesan...")

	// Listen forever
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("📥 Geofence Event Diterima: %s\n", string(msg.Body))
		}
	}()

	<-forever
}
