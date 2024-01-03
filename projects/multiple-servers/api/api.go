package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Run(port int, dbUrl string) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("unable to connect to db: %v\n", err)
	}
	defer conn.Close(context.Background())

	http.HandleFunc("/api/images.json", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.EscapedPath())
		w.Header().Add("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case http.MethodGet:
			getImages(w, r, conn)
		case http.MethodPost:
			postImages(w, r, conn)
		default:
			http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Listening on http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func getImages(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	images, err := fetchImages(conn, r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	respondWithIndentedData(w, r, images)
}

func postImages(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var image Image
	err = json.Unmarshal(body, &image)
	if err != nil {
		http.Error(w, "error unmarshaling the body", http.StatusBadRequest)
		return
	}

	width, height, err := fetchImageResolution(image.URL)
	if err != nil {
		http.Error(w, "no valid image provided", http.StatusBadRequest)
		return
	}
	image.Width = float64(width)
	image.Height = float64(height)

	err = saveImage(conn, r.Context(), image)
	if err != nil {
		http.Error(w, "error saving the image to the database", http.StatusInternalServerError)
		return
	}

	respondWithIndentedData(w, r, image)
}

func respondWithIndentedData(w http.ResponseWriter, r *http.Request, data any) {
	var b []byte
	var err error

	identParam := r.URL.Query().Get("indent")
	if identParam == "" {
		b, err = json.Marshal(data)
	} else {
		indent, convErr := strconv.Atoi(identParam)
		if convErr != nil {
			http.Error(w, "indent has to be a number", http.StatusBadRequest)
			return
		}
		b, err = json.MarshalIndent(data, "", strings.Repeat(" ", indent))
	}

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
