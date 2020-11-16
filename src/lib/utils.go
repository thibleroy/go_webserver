package lib

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func dotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetServerEnv() IEnvironment {
	wspStr := dotEnvVariable("WEBSERVER_PORT")
	mongoURLStr := dotEnvVariable("MONGODB_URL")
	mongoPortStr := dotEnvVariable("MONGODB_PORT")
	var serverEnvironment IEnvironment
	wsPort, _ := strconv.Atoi(wspStr)
	mongoPort, _ := strconv.Atoi(mongoPortStr)
	serverEnvironment = IEnvironment{WebServerPort: wsPort, MongoPort: mongoPort, MongoURL: mongoURLStr}
	return serverEnvironment
}
