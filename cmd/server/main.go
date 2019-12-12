package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/rodrwan/twoppin/repository/database"
	"github.com/rodrwan/twoppin/services/graphql"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := database.NewPostgres("postgres://twoppin:twoppin2020@localhost:5432/twoppin?sslmode=disable")
	check(err)

	rootHandler := handler.GraphQL(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: &graphql.Resolver{
					Database: db,
				},
			},
		),
	)

	port := 8000
	mux := http.NewServeMux()
	mux.Handle("/", handler.Playground("GraphQL playground", "/query"))
	mux.Handle("/query", rootHandler)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(server.ListenAndServe())
}
