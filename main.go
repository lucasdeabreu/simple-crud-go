package main

import (
	"fmt"
	"log"
	"net/http"

	infra "github.com/lucasdeabreu/simple-crud-go/infra"
	"github.com/lucasdeabreu/simple-crud-go/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	handler := user.Handler{Service: &user.Service{Db: infra.Db}}
	router := httprouter.New()

	router.GET("/api/v1/users", handler.GetAll)
	router.GET("/api/v1/users/:id", handler.GetByID)
	router.POST("/api/v1/users", handler.Create)
	router.PUT("/api/v1/users/:id", handler.Update)
	router.DELETE("/api/v1/users/:id", handler.DeleteByID)

	fmt.Println("Starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
