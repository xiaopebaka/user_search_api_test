package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"my_gql_server/graph"
	dataType "my_gql_server/graph/db"
	"my_gql_server/graph/services"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Automatically migrate schema
	err = db.AutoMigrate(&dataType.Users{}, &dataType.Locations{}, &dataType.Icons{}, &dataType.UserRegistrationStatus{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	service := services.New(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Srv: service,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
