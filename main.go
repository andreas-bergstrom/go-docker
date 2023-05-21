package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonData := map[string]string{
			"hello": "world",
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(jsonData)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	fmt.Printf("Server running (port=%s), route: http://localhost:%s/\n", port, port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
