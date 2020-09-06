package api

import (
	"fmt"
	"log"
	"os"

	"github.com/AlexSwiss/reev-v1/api/controllers"
	"github.com/AlexSwiss/reev-v1/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Run initializes server and starts the api on port 8080
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")

}
