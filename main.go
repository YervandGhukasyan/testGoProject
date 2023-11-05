package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	// "github.com/go-chi/core"
)

func main() {
	fmt.Println("Voghjuyn Yerr")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not a found in the enviroment")
	}

	router := chi.NewRouter()

	// router.Use(
	// 	core.Handle()
	// )

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server runing on port %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT: ", portString)
}
