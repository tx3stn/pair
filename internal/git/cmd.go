// Package git contains logic for git operations.
package git

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Commit commits the currently staged files with the msg (generated from the
// pairing values and free text input) and the args specified in the config file.
func Commit(ctx context.Context, msg string, args string) error {
	cmdArgs := []string{"commit"}

	if strings.TrimSpace(args) != "" {
		cmdArgs = append(cmdArgs, strings.Fields(args)...)
	}

	cmdArgs = append(cmdArgs, "-m", msg)

	return gitCommand(ctx, cmdArgs...)
}

func gitCommand(ctx context.Context, args ...string) error {
	// #nosec G204 -- args are intentional git CLI flags/subcommands
	cmd := exec.CommandContext(ctx, "git", args...)

	// Inherit the terminal streams so git (and any commit hooks) write their
	// output directly to the user's terminal instead of being swallowed.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %w", ErrRunningGitCommand, err)
	}

	return nil
}
