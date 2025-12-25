package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello");

	godotenv.Load();
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the enviornment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"}, // Specify allowed origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Include OPTIONS for preflight
		AllowedHeaders: []string{"*"}, // Specify allowed headers
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum duration (in seconds) the results of a preflight request can be cached
	}))

	v1Router := chi.NewRouter();
	v1Router.Get("/healthz", handlerReadiness);
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router);

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString ,
	}

	log.Printf("Server Starting on port %v", portString);
	err := srv.ListenAndServe();

	if err != nil {
		log.Fatal(err)
	}

}