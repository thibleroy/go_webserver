package lib

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func dotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetServerEnv()(error, IEnvironment) {
	var strvar = dotEnvVariable("WEBSERVER_PORT")
	var serverEnvironment IEnvironment
	i, err := strconv.Atoi(strvar)
	if err != nil {
		fmt.Println("get server env error", err)
	} else {
		serverEnvironment = IEnvironment{Port: i}
		fmt.Println("get server env success:", serverEnvironment)
	}
	return err, serverEnvironment
}
