// Package main is the entrypoint of the CLI.
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/tx3stn/pair/cmd"
)

func main() {
	code := 0

	defer func() {
		os.Exit(code)
	}()

	ctx := context.Background()

	app := cmd.NewApp()
	if err := app.ExecuteContext(ctx); err != nil {
		code = 1

		slog.Error(err.Error())
	}
}
