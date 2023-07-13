package main

import (
	"fmt"
	"github.com/D3FiX4M/go-restfull/initializers"
	"github.com/D3FiX4M/go-restfull/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal("Migration failed", err)
	}
	fmt.Println("? Migration complete")
}
