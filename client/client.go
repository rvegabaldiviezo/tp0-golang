package main

import (
	"fmt"
	"log"
	"log/slog"

	//"strings"
	globals "github.com/rvegabaldiviezo/tp0-golang/client/globals"

	utils "github.com/rvegabaldiviezo/tp0-golang/shared/utils"
)

func main() {

	/** Inicializamos las configs.
	 */
	globals.ClientConfig = utils.IniciarConfiguracion[globals.Config]()
	if globals.ClientConfig == nil {
		log.Fatal("Error en la lectura de la config.json")
	}

	// Inicilializamos los logs.
	logFile := utils.ConfigurarLogger()
	utils.SetLevelLog(globals.ClientConfig.LogLevel)
	// defer: pospone la ejecución de logFile.Close() hasta que la función actual termine.
	defer logFile.Close()

	// loggeamos el valor de la config
	slog.Debug("Puerto: " + fmt.Sprintf("%v", globals.ClientConfig.Puerto))
	slog.Info("Mensaje: " + globals.ClientConfig.Mensaje)
	slog.Warn("LogLevel: " + globals.ClientConfig.LogLevel)
	slog.Error("IP: " + globals.ClientConfig.Ip)

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
