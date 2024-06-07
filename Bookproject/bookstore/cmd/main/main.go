package main

import (
	"log"
	"net/http"

	"bookstore/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
