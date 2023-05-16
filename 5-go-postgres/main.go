package main

import (
	"fmt"
	"log"
	"net/http"
	"postgresql/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
