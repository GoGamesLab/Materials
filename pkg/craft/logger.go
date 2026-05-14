package craft

import (
	"flag"
	"log/slog"
	"os"

	"github.com/GoGamesLab/Materials/pkg/config"
)

var Logger *slog.Logger

func translateLogLevel(s string) slog.Level {
	switch s {
	case "Debug":
		return slog.LevelDebug
	case "Info":
		return slog.LevelInfo
	case "Warn", "Warning":
		return slog.LevelWarn
	case "Error":
		return slog.LevelError
	default:
		panic("Invalid log level")
	}
}

func init() {
	configDir := flag.String("configDir", "config", "Caminho para pasta de configuração")
	flag.Parse()

	c := config.Load(*configDir)

	Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(translateLogLevel(c.Application.Log.Level))}))
}
