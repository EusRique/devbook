package main

import (
	"devbook/src/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	println(config.StringConexaoBanco)
	println("API Rodando!!!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
