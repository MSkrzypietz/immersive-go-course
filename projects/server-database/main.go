package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Image struct {
	Title   string  `json:"title"`
	AltText string  `json:"alt_text"`
	URL     string  `json:"url"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
}

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_URL not provided")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to db: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	http.HandleFunc("/images.json", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getImages(w, r, conn)
		case http.MethodPost:
			postImages(w, r, conn)
		default:
			http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func getImages(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	rows, err := conn.Query(r.Context(), "SELECT title, url, alt_text, width, height FROM public.images")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var title, url, altText string
		var width, height int
		err = rows.Scan(&title, &url, &altText, &width, &height)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		images = append(images, Image{Title: title, AltText: altText, URL: url, Width: float64(width), Height: float64(height)})
	}

	respondWithIndentedData(w, r, images)
}

func postImages(w http.ResponseWriter, r *http.Request, conn *pgx.Conn) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusBadRequest)
		return
	}

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

	rows, err := conn.Query(r.Context(), "INSERT INTO public.images (title, url, alt_text, width, height) VALUES ($1, $2, $3, $4, $5)", image.Title, image.URL, image.AltText, width, height)
	if err != nil {
		http.Error(w, "error storing the image to the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

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

func fetchImageResolution(url string) (width int, height int, err error) {
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("image request failed: %w", err)
	}
	defer resp.Body.Close()

	config, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("image decoding failed: %w", err)
	}
	return config.Width, config.Height, nil
}
