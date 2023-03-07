package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
	"http-client/testUser/cmd/api/di"
	"log"
	"net/http"
)

func main() {
	mux, err := di.Initialize()
	if err != nil {
		log.Fatalf("Fallo de la construccion de los servicios: %v", err)
	}

	err = http.ListenAndServe(":8080", mux)
	fmt.Println(err.Error())
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Cerrando servidor")
	} else if err != nil {
		log.Println("Error del servidor por: " + err.Error())

	}
	log.Println("servidor corriendo")

}
