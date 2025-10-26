package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/service_transjakarta/api_golang_tj/database"
	"example.com/service_transjakarta/api_golang_tj/database/rabbit"
	"example.com/service_transjakarta/api_golang_tj/routes"
)

func main() {
	fmt.Println("Aplikasi GO sedang berjalan")

	fileActivity, err := os.OpenFile("logs/activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(fileActivity)

	db, err := database.Connect()
	if err != nil {
		fmt.Println("❌ Gagal koneksi DB:", err)
		return
	}
	defer db.Close()

	rabbit.ConnectRabbit()

	clientMqtt := database.ConnectMqtt()
	go database.SubscribeMqtt(clientMqtt, db)
	database.PublishMockLocation(clientMqtt)

	router := routes.InitRoutes(db)
	http.Handle("/", router)

	fmt.Println("Server berjalan di port 7002")
	log.Fatal(http.ListenAndServe(":7002", nil))
}
