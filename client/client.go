package main

import (
	"log"
	//"strings"
	clientglobals "github.com/rvegabaldiviezo/tp0-golang/client/globals"
	utils "github.com/rvegabaldiviezo/tp0-golang/shared/utils"
)

func main() {

	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Vamos a promocionar operativos")

	clientglobals.ClientConfig = utils.IniciarConfiguracion[clientglobals.Config]("configs/config.json")
	// validar que la config este cargada correctamente
	if clientglobals.ClientConfig == nil {
		log.Fatal("Error en la lectura de la config.json")
	}

	// loggeamos el valor de la config
	log.Printf("%+v\n", clientglobals.ClientConfig)
	log.Println("Mensaje: ", clientglobals.ClientConfig.Mensaje)
	log.Println("IP: ", clientglobals.ClientConfig.Ip)
	log.Println("Puerto: ", clientglobals.ClientConfig.Puerto)
	log.Println("LogLevel: ", clientglobals.ClientConfig.LogLevel)

	// // ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// // enviar un mensaje al servidor con el valor de la config

	// // leer de la consola el mensaje

	// for {
	// 	input := utils.LeerConsola()
	// 	log.Printf("Texto recibido: %q\n", input)
	// 	if strings.TrimSpace(input) == "" {
	// 		log.Println("Se ingreso un salto de linea, finalizamos la lectura por consola")
	// 		break
	// 	}
	// }

	// generamos un paquete y lo enviamos al servidor

	// utils.GenerarYEnviarPaquete()
}
