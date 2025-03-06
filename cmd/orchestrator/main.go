package main

import (
	"log"

	"github.com/TravellerGSF/distributed-calculator/internal/orchestrator/api"
)

func main() {
	log.Println("Запуск оркестратора...")
	if err := api.SetupRouter(); err != nil {
		log.Fatalf("Ошибка при запуске оркестратора: %v", err)
	}
}
