package main

import (
	"log"
	"net/http"

	serverutils "github.com/rvegabaldiviezo/tp0-golang/shared/connections"

	utils "github.com/rvegabaldiviezo/tp0-golang/shared/utils"
)

func main() {

	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Vamos a levantar el server")

	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", serverutils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", serverutils.RecibirMensaje)

	//panic("no implementado!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
