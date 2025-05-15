package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"

	"hackathons/config"
	"hackathons/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfgPath := flag.String("c", "config/config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.ReadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
