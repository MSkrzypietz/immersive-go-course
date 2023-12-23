package main

import (
	"bytes"
	"fmt"
	"golang.org/x/time/rate"
	"html"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		if r.Method == http.MethodPost {
			data, _ := io.ReadAll(r.Body)
			defer r.Body.Close()
			w.Write([]byte(html.EscapeString(string(data))))
			return
		}

		var out bytes.Buffer
		out.WriteString("<!DOCTYPE html><html><em>Hello, World</em><p>Query parameters:</p><ul>")
		for key, values := range r.URL.Query() {
			out.WriteString(fmt.Sprintf("<li>%s: [%s]</li>", html.EscapeString(key), html.EscapeString(strings.Join(values, " "))))
		}
		out.WriteString("</ul></html>")
		w.Write(out.Bytes())
	})

	http.HandleFunc("/200", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("200"))
	})

	http.Handle("/404", http.NotFoundHandler())

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	})

	authUsername := os.Getenv("AUTH_USERNAME")
	authPassword := os.Getenv("AUTH_PASSWORD")
	if authUsername == "" || authPassword == "" {
		panic("Auth username or password is not set")
	}
	http.HandleFunc("/authenticated", func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || authUsername != username || authPassword != password {
			w.Header().Add("WWW-Authenticate", "Basic realm=\"localhost\", charset=\"UTF-8\"")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})

	limiter := rate.NewLimiter(100, 30)
	http.HandleFunc("/limited", func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many requests"))
			return
		}

		w.Write([]byte("Success response"))
	})

	http.ListenAndServe(":8080", nil)
}
