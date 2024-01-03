package static

import (
	"fmt"
	"log"
	"net/http"
)

func Run(port int, path string) {
	fsHandler := http.FileServer(http.Dir(path))
	log.Printf("Listening on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), fsHandler)
	if err != nil {
		log.Fatalln(err)
	}
}
