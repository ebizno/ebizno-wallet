package main

import (
	"log"
	"path/filepath"
	"runtime"

	RestServerWellet "github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/server"
	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func main() {
	RestServerWellet.WelletServerRest()
}
