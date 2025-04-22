package utils

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// Nivel mínimo de logs a mostrar. Los niveles en orden de mayor a menor severidad son:
// Niveles de mayor a menor:  ERROR > WARN > INFO > DEBUG
//  1. ERROR (más alto): solo se muestran errores críticos
//  2. WARN             : advertencias (no detienen el programa, pero son relevantes)
//  3. INFO             : mensajes informativos (estado general del sistema)
//  4. DEBUG (más bajo) : información detallada para desarrolladores, como valores de variables, etc.
//
// IMPORTANTE:
// Establecer un nivel más ALTO implica que los logs de nivel más BAJO serán ignorados.
// Por ejemplo:
//   - Si usás Level=INFO, no se verán los logs DEBUG.
//   - Si usás Level=ERROR, solo aparecerán los errores (no INFO, WARN ni DEBUG).
//   - Si usás Level=DEBUG, vas a ver todos los logs.
var NivelesLogs = map[string]slog.Level{
	"ERROR": slog.LevelError,
	"WARN":  slog.LevelWarn,
	"INFO":  slog.LevelInfo,
	"DEBUG": slog.LevelDebug,
}

var logLevelVar = new(slog.LevelVar) // Creamos un puntero a slog.LevelVar. Nivel dinámico (por defecto es Info)

// ConfigurarLogger: inicializa el logger para el módulo desde el cual se llama a esta funcion.
func ConfigurarLogger() *os.File {

	// Seteamos el nivel por defecto al iniciar el paquete
	logLevelVar.Set(slog.LevelDebug)

	// Armar el path del nombre del Archivo.
	path := filepath.Join("logs", GenerarNombreArchivoLog())

	// Crear archivo de log
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	// No cerramos acá, lo deberías cerrar al final del main()
	//defer logFile.Close() // Se usa desde el main

	// Crear MultiWriter (terminal + archivo)
	mw := io.MultiWriter(os.Stdout, logFile)

	// Crear el handler JSON con nivel mínimo Debug
	handler := slog.NewJSONHandler(mw, &slog.HandlerOptions{
		Level: logLevelVar,
	})

	// Setear como logger por defecto
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("Inicializo los logs por terminal y por archivo", "path: ", path)
	// retornamos para poder cerrarlo mas adelante
	return logFile
}

func SetLevelLog(level string) {
	if lvl, ok := NivelesLogs[level]; ok {
		logLevelVar.Set(lvl)
		slog.Info("Nivel de log actualizado", "nivel", level)
	} else {
		slog.Warn("Nivel de log inválido. Usar uno de: ERROR, WARN, INFO, DEBUG", "nivelProporcionado", level)
	}
}
