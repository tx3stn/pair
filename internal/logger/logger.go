// Package logger configures the logger.
// Used so additional information can be output when running in verbose mode.
package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/fatih/color"
	"github.com/tx3stn/pair/internal/flags"
)

func New(accessible bool) {
	level := slog.LevelInfo
	if flags.Verbose {
		level = slog.LevelDebug
	}

	handler := &customHandler{w: os.Stderr, level: level, accessible: accessible}
	slog.SetDefault(slog.New(handler))
}

type customHandler struct {
	w          io.Writer
	level      slog.Level
	accessible bool
}

func (h *customHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *customHandler) Handle(_ context.Context, r slog.Record) error {
	args := ""
	keyColor := getKeyColor(r.Level)

	r.Attrs(func(a slog.Attr) bool {
		if args != "" {
			args += " "
		}

		if h.accessible {
			args += fmt.Sprintf("%s=%v", a.Key, a.Value)
		} else {
			args += fmt.Sprintf("%s=%v", keyColor.Sprint(a.Key), a.Value)
		}

		return true
	})

	if args != "" {
		args = " " + args
	}

	levelStr := r.Level.String()
	if !h.accessible {
		levelStr = colorizeLevel(r.Level)
	}

	_, err := fmt.Fprintf(h.w, "%s %s%s\n", levelStr, r.Message, args)

	return fmt.Errorf("error handling slog message: %w", err)
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return h
}

func colorizeLevel(level slog.Level) string {
	return getKeyColor(level).Sprint(fmt.Sprintf("🍐 %s:", getLevelText(level)))
}

func getLevelText(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return "DEBUG"
	case slog.LevelInfo:
		return "INFO"
	case slog.LevelWarn:
		return "WARN"
	case slog.LevelError:
		return "ERROR"
	default:
		return level.String()
	}
}

func getKeyColor(level slog.Level) *color.Color {
	switch level {
	case slog.LevelDebug:
		return color.New(color.FgCyan)
	case slog.LevelInfo:
		return color.New(color.FgGreen)
	case slog.LevelWarn:
		return color.New(color.FgYellow)
	case slog.LevelError:
		return color.New(color.FgRed)
	default:
		return color.New()
	}
}
