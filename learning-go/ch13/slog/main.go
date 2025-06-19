package main

import "log/slog"

func main() {
	slog.Debug("debug log message")
	slog.Warn("warn log message")
	slog.Info("info log message")
	slog.Error("error log message")
}
