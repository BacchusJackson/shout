package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

var (
	cliVersion     = "[Not Provided]"
	buildDate      = "[Not Provided]"
	gitCommit      = "[Not Provided]"
	gitDescription = "[Not Provided]"
)

func main() {
	Run()
}

func Run() int {
	h := tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
	})
	slog.SetDefault(slog.New(h))

	if err := NewCommand().Execute(); err != nil {
		return -1
	}

	return 0
}
