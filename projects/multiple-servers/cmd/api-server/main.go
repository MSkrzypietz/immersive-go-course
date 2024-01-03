package main

import (
	"flag"
	"log"
	"multiple-servers/api"
	"os"
)

func main() {
	port := flag.Int("port", 8081, "port to listen on")
	flag.Parse()

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatalln("You have to provide a database connection url with DATABASE_URL")
	}

	api.Run(*port, dbUrl)
}
