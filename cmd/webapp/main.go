package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sheghun/containerized-webapp/internal/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.RenderWebPage)
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./web/build")))

	r.HandleFunc("/api", handler.FindHighestPrime).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedMethods:   []string{http.MethodOptions, http.MethodPost, http.MethodGet},
	})

	// start server listen
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8080"
	}

	log.Println("Server listening on ", PORT)
	if err := http.ListenAndServe(PORT, c.Handler(r)); err != nil {
		log.Fatalf("error occured starting server: %v", err)
	}
}
