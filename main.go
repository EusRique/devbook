package main

import (
	"devbook/src/router"
	"log"
	"net/http"
)

func main() {
	println("Hello Worlds!!!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":8080", r))
}
