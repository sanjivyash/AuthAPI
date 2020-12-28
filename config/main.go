package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load("config/dev.env")

	if err != nil {
		log.Fatal("Error loading .env file\n" + err.Error())
	}

	val, ok := os.LookupEnv(key)

	if !ok {
		fmt.Println("[ERROR] Invalid key: " + key)
	}

	if key == "BASE_DIR" {
		if gopath, ok := os.LookupEnv("GOPATH"); ok {
			val = gopath + val 
		} else {
			log.Fatal("GOPATH not set as environment variable")
		}
	}

	return val
}
