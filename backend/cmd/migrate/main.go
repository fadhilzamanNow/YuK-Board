package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"yukboard/config"
)

func main() {
	fresh := flag.Bool("fresh", false, "Drop all tables before migrating")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()

	if *fresh {
		config.DropAll()
	}

	config.Migrate()
}
