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
	// e.g.: git commit -S -m "feat(JIRA-100): did this thang"
	return gitCommand(ctx, "commit", args, "-m", msg)
}

func gitCommand(ctx context.Context, args ...string) (string, error) {
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
