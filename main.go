package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SameerNadeem/rssag/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL is not found")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("database connection failed:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://k", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Get("/users",apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("server starting point %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port:", portString)
}
