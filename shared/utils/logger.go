package utils

import (
	"io"
	"log"
	"log/slog"
	"os"
)

// Mapa de strings a slog.Level
var NivelesLogs = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

/*
ConfigurarLogger: inicializa el logger para el módulo indicado.

	module string: recibe el nombre del modulo
*/
// func ConfigurarLogger() {
// 	module := GenerarNombreArchivoLog()
// 	logFile, err := os.OpenFile("logs/"+module, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	mw := io.MultiWriter(os.Stdout, logFile)
// 	log.SetOutput(mw)
// }

func ConfigurarLogger() {
	module := GenerarNombreArchivoLog()

	logFile, err := os.OpenFile("logs/"+module, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close() // Se usa desde el main

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}

// func ConfigurarLogger() {
// 	module := GenerarNombreArchivoLog()

// 	path := filepath.Join("logs", module)
// 	fmt.Println("path: ", path)

// 	// Crear archivo de log
// 	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// No cerramos acá, lo deberías cerrar al final del main()

// 	// Crear MultiWriter (terminal + archivo)
// 	mw := io.MultiWriter(os.Stdout, logFile)

// 	// Crear el handler JSON con nivel mínimo Debug
// 	handler := slog.NewJSONHandler(mw, &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	})

// 	// Setear como logger por defecto
// 	logger := slog.New(handler)
// 	slog.SetDefault(logger)
// }
