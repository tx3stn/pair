// Package pairing contains logic around current pairing session.
package pairing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tx3stn/pair/internal/git"
)

const DataDir = "/tmp/pair"

type Session struct {
	parentDir  string
	on         string
	with       []git.CoAuthor
	SessionDir string
	WithFile   string
	OnFile     string
}

func NewSession(dir string) *Session {
	return &Session{
		parentDir:  dir,
		on:         "",
		with:       []git.CoAuthor{},
		SessionDir: dir,
		WithFile:   filepath.Join(dir, "with"),
		OnFile:     filepath.Join(dir, "on"),
	}
}

func (s *Session) Clean() error {
	err := os.Remove(s.WithFile)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error removing file %s: %w", s.WithFile, err)
	}

	err = os.Remove(s.OnFile)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error removing file %s: %w", s.OnFile, err)
	}

	return nil
}

func (s *Session) GetCoAuthors() ([]git.CoAuthor, error) {
	data, err := os.ReadFile(s.WithFile)
	if os.IsNotExist(err) {
		return []git.CoAuthor{}, nil
	}

	if err != nil {
		return []git.CoAuthor{}, fmt.Errorf("%w: %w", ErrReadingCoAuthors, err)
	}

	lines := bytes.Split(data, []byte{'\n'})
	coAuthors := make([]git.CoAuthor, 0, len(lines))

	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		var coAuthor git.CoAuthor
		if err := json.Unmarshal(line, &coAuthor); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrUnmarshalingCoAuthor, err)
		}

		coAuthors = append(coAuthors, coAuthor)
	}

	s.with = coAuthors

	return s.with, nil
}

func (s *Session) SetCoAuthors(coAuthors []git.CoAuthor) error {
	s.with = coAuthors

	if err := os.MkdirAll(s.SessionDir, 0o750); err != nil {
		return fmt.Errorf("%w: %w", ErrCreatingDirectory, err)
	}

	var buf bytes.Buffer

	for i, coAuthor := range s.with {
		b, err := json.Marshal(coAuthor)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrMarshalingCoAuthor, err)
		}

		buf.Write(b)

		if i < len(coAuthors)-1 {
			buf.WriteByte('\n')
		}
	}

	if err := os.WriteFile(s.WithFile, buf.Bytes(), 0o600); err != nil {
		return fmt.Errorf("%w: %w", ErrWritingCoAuthors, err)
	}

	return nil
}

func (s *Session) GetTicketID() (string, error) {
	data, err := os.ReadFile(s.OnFile)
	if os.IsNotExist(err) {
		return "", nil
	}

	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrReadingTicketID, err)
	}

	s.on = strings.TrimSpace(string(data))

	return s.on, nil
}

func (s *Session) SetTicketID(ticketID string) error {
	s.on = ticketID

	if err := os.MkdirAll(s.SessionDir, 0o750); err != nil {
		return fmt.Errorf("%w: %w", ErrCreatingDirectory, err)
	}

	if err := os.WriteFile(s.OnFile, []byte(s.on), 0o600); err != nil {
		return fmt.Errorf("%w: %w", ErrWritingTicketID, err)
	}

	return nil
}
