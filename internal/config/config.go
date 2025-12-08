// Package config loads env fils into a Login struct which gets passed to the rest of the application
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Login struct {
	User string
	Pass string
	IP   string
	Port string
	Name string
}

func CreateEnvFile() error {
	err := os.WriteFile(".env", nil, 0o666)
	if err != nil {
		return err
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv("DATABASE_USER", "nasr")
	os.Setenv("DATABASE_PASSWORD", "root")
	os.Setenv("DATABASE_IP", "127.0.0.1")
	os.Setenv("DATABASE_PORT", "3306")
	os.Setenv("DATABASE_NAME", "synf")

	return nil
}

func (L *Login) LoadCredentials() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error in loading credentials")
	}
	L.User = os.Getenv("DATABASE_USER")
	L.Pass = os.Getenv("DATABASE_PASSWORD")
	L.IP = os.Getenv("DATABASE_IP")
	L.Port = os.Getenv("DATABASE_PORT")
	L.Name = os.Getenv("DATABASE_NAME")
}
