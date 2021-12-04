package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Port int = 3000
)

func Load() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 3000
	}

	//message := fmt.Sprintf("VAR_NAME: NAME AND VALUE: %s", os.Getenv("NAME"))
	//fmt.Println(message)
}
