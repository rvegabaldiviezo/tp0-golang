package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func LeerConsola() string {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")
	text, _ := reader.ReadString('\n')
	log.Print(text)
	return text
}

func NombreDelModulo() string {
	// Obtiene el path absoluto del directorio actual
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error al obtener el directorio actual:", err)
		return "module"
	}
	// Extrae solo el nombre del directorio (basename)
	return filepath.Base(wd)
}

// GenerarNombreArchivoLog genera un nombre de archivo de log con el formato:
// {timestamp}_{id_aleatorio}_{nombreModulo}.log
func GenerarNombreArchivoLog() string {
	// Extraído del nombre del directorio actual
	nombreModulo := NombreDelModulo()
	// Zona horaria de Argentina
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		loc = time.FixedZone("UTC-3", -3*60*60) // fallback si falla
	}
	// Timestamp con milisegundos en formato ISO con zona horaria argentina (UTC-3)
	timestamp := time.Now().In(loc).Format("2006-01-02T15-04-05T07")

	// ID aleatorio de 6 dígitos
	// id := rand.Intn(900000) + 100000

	// Nombre del archivo
	//return fmt.Sprintf("%s_%d_%s.log", timestamp, id, nombreModulo)
	return fmt.Sprintf("%s_%s.log", timestamp, nombreModulo)
}
