// Package git contains logic for git operations.
package git

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Commit commits the currently staged files with the msg (generated from the
// pairing values and free text input) and the args specified in the config file.
func Commit(ctx context.Context, msg string, args string) (string, error) {
	cmdArgs := []string{"commit"}

	if strings.TrimSpace(args) != "" {
		cmdArgs = append(cmdArgs, strings.Fields(args)...)
	}

	cmdArgs = append(cmdArgs, "-m", msg)

	return gitCommand(ctx, cmdArgs...)
}

func gitCommand(ctx context.Context, args ...string) (string, error) {
	// #nosec G204 -- args are intentional git CLI flags/subcommands
	cmd := exec.CommandContext(ctx, "git", args...)

	var stdOut bytes.Buffer

	var stdErr bytes.Buffer

	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrRunningGitCommand, err)
	}

	return strings.Trim(stdOut.String(), "\n"), nil
}
