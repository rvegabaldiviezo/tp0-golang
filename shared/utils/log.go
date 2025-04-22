// Package logger proporciona funciones simplificadas para loguear con slog
// y manejar errores fatales de forma más elegante y centralizada.
package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

// CustomHandler imprime logs con colores en consola según el nivel de log
// Niveles de mayor a menor:
//   ERROR > WARN > INFO > DEBUG
// Nivel más bajo (DEBUG): muestra todos los logs
// Nivel más alto (ERROR): muestra solo errores
// Esto implica que un nivel más alto excluye mensajes de menor severidad

type CustomHandler struct {
	level slog.Level
}

func (h *CustomHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *CustomHandler) Handle(_ context.Context, r slog.Record) error {
	color := nivelAColor(r.Level)
	reset := "\033[0m"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := r.Message

	fmt.Fprintf(os.Stdout, "%s[%s] [%s] %s%s\n",
		color,
		timestamp,
		r.Level.String(),
		msg,
		reset,
	)

	return nil
}

func (h *CustomHandler) WithAttrs(_ []slog.Attr) slog.Handler { return h }
func (h *CustomHandler) WithGroup(_ string) slog.Handler      { return h }

func nivelAColor(level slog.Level) string {
	switch level {
	case slog.LevelError:
		return "\033[31m" // rojo
	case slog.LevelWarn:
		return "\033[33m" // amarillo
	case slog.LevelInfo:
		return "\033[36m" // cian
	case slog.LevelDebug:
		return "\033[37m" // gris claro
	default:
		return "\033[0m"
	}
}

// InitLogger inicializa slog con el CustomHandler coloreado para consola
func InitLogger() {
	handler := &CustomHandler{
		level: slog.LevelDebug,
	}
	slog.SetDefault(slog.New(handler))
}

// Funciones auxiliares de logging

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}
