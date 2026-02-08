// Package main is the entrypoint of the CLI.
package main

import (
	"context"
	"fmt"
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

		fmt.Printf("%s\n", err.Error())
	}
}
