package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}

	return os.Getenv("MONGO_URI")
}

func GetEnvDatabaseName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}

	return os.Getenv("DATABASE_NAME")
}

func GetEnvPortNumber() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error .env")
	}

	return os.Getenv("PORT_NUMBER")
}
