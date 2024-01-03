package api

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"image"
	"net/http"
	"time"
)

type Image struct {
	Title   string  `json:"title"`
	AltText string  `json:"alt_text"`
	URL     string  `json:"url"`
	Width   float64 `json:"width"`
	Height  float64 `json:"height"`
}

func fetchImages(conn *pgx.Conn, ctx context.Context) ([]Image, error) {
	var images []Image

	rows, err := conn.Query(ctx, "SELECT title, url, alt_text, width, height FROM public.images")
	if err != nil {
		return images, fmt.Errorf("error quering the db: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var title, url, altText string
		var width, height int
		err = rows.Scan(&title, &url, &altText, &width, &height)
		if err != nil {
			return images, fmt.Errorf("error interpreting the db repsonse: %w", err)
		}
		images = append(images, Image{Title: title, AltText: altText, URL: url, Width: float64(width), Height: float64(height)})
	}

	return images, nil
}

func saveImage(conn *pgx.Conn, ctx context.Context, image Image) error {
	rows, err := conn.Query(ctx, "INSERT INTO public.images (title, url, alt_text, width, height) VALUES ($1, $2, $3, $4, $5)", image.Title, image.URL, image.AltText, image.Width, image.Height)
	if err != nil {
		return fmt.Errorf("error saving the image to the database: %w", err)
	}
	defer rows.Close()
	return nil
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
