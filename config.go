package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       "8080",
		DBUser:     "root",
		DBPassword: "", 
		DBAddress:  fmt.Sprintf("127.0.0.1:3306"),
		DBName:     "projectmanager",
		JWTSecret:  "randomjwtsecretkey",
	}
}


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}