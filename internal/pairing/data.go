// Package pairing contains logic around current pairing session.
package pairing

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tx3stn/pair/internal/git"
)

const DataDir = "/tmp/pair"

type Session struct {
	date      time.Time
	parentDir string
	on        string
	with      []string
}

func NewSession(dir string, currDate time.Time) Session {
	return Session{
		date:      currDate,
		parentDir: dir,
		on:        "",
		with:      []string{},
	}
}

func (s Session) Clean() error {
	if err := os.RemoveAll(s.activeDir()); err != nil {
		return fmt.Errorf("error removing directory %s: %w", s.activeDir(), err)
	}

	return nil
}

func (s Session) GetCoAuthors() ([]string, error) {
	data, err := os.ReadFile(s.GetPath("with"))
	if os.IsNotExist(err) {
		return []string{}, nil
	}

	if err != nil {
		return []string{}, fmt.Errorf("%w: %w", ErrReadingCoAuthors, err)
	}

	content := strings.TrimSpace(string(data))
	if content == "" {
		return []string{}, nil
	}

	s.with = strings.Split(content, "\n")

	return s.with, nil
}

func (s Session) SetCoAuthors(coAuthors []git.CoAuthor) ([]string, error) {
	formatted := make([]string, len(coAuthors))
	for i, author := range coAuthors {
		formatted[i] = author.Format()
	}

	s.with = formatted

	if err := os.MkdirAll(s.activeDir(), 0o750); err != nil {
		return []string{}, fmt.Errorf("%w: %w", ErrCreatingDirectory, err)
	}

	stringOutput := strings.Join(s.with, "\n")
	if err := os.WriteFile(s.GetPath("with"), []byte(stringOutput), 0o600); err != nil {
		return []string{}, fmt.Errorf("%w: %w", ErrWritingCoAuthors, err)
	}

	return formatted, nil
}

func (s Session) GetTicketID() (string, error) {
	data, err := os.ReadFile(s.GetPath("on"))
	if os.IsNotExist(err) {
		return "", nil
	}

	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrReadingTicketID, err)
	}

	s.on = strings.TrimSpace(string(data))

	return s.on, nil
}

func (s Session) SetTicketID(ticketID string) error {
	s.on = ticketID

	if err := os.MkdirAll(s.activeDir(), 0o750); err != nil {
		return fmt.Errorf("%w: %w", ErrCreatingDirectory, err)
	}

	if err := os.WriteFile(s.GetPath("on"), []byte(s.on), 0o600); err != nil {
		return fmt.Errorf("%w: %w", ErrWritingTicketID, err)
	}

	return nil
}

func (s Session) GetPath(fileName string) string {
	return filepath.Join(s.activeDir(), fileName)
}

func (s Session) activeDir() string {
	return filepath.Join(s.parentDir, s.date.Format("2006-01-02"))
}
