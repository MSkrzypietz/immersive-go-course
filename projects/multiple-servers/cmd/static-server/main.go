package main

import (
	"flag"
	"log"
	"multiple-servers/static"
)

func main() {
	port := flag.Int("port", 8082, "port to listen on")
	path := flag.String("path", "", "path to static files")
	flag.Parse()

	if *path == "" {
		log.Fatalln("You need to specify the path to the static files")
	}

	static.Run(*port, *path)
}
