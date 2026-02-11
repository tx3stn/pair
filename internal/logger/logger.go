// Package logger configures the logger.
// Used so additional information can be output when running in verbose mode.
package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

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

	r.Attrs(func(a slog.Attr) bool {
		if args != "" {
			args += " "
		}

		args += fmt.Sprintf("%s=%v", a.Key, a.Value)

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
	icon := "🍐 "

	switch level {
	case slog.LevelDebug:
		return fmt.Sprintf("\033[36m%sDEBUG:\033[0m", icon)
	case slog.LevelInfo:
		return fmt.Sprintf("\033[32m%s INFO:\033[0m", icon)
	case slog.LevelWarn:
		return fmt.Sprintf("\033[33m%s  WARN:\033[0m", icon)
	case slog.LevelError:
		return fmt.Sprintf("\033[31m%s ERROR:\033[0m", icon)
	default:
		return level.String()
	}
}
