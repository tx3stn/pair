package git

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode/utf8"
)

type CoAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c CoAuthor) Format() string {
	return fmt.Sprintf("Co-authored-by: %s <%s>", c.Name, c.Email)
}

type emailTemplateData struct {
	FirstName    string
	LastName     string
	FirstInitial string
	LastInitial  string
}

// SuggestEmail renders the given template format for a co-author's email based
// on their first and last name values.
//
// The template has access to the fields: FirstName, LastName, FirstInitial and
// LastInitial, e.g. "{{.FirstInitial}}{{.LastName}}@example.com".
func SuggestEmail(format string, firstName string, lastName string) (string, error) {
	if format == "" {
		return "", nil
	}

	tmpl, err := template.New("email").Option("missingkey=error").Parse(format)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrSuggestingEmail, err)
	}

	data := emailTemplateData{
		FirstName:    firstName,
		LastName:     lastName,
		FirstInitial: initial(firstName),
		LastInitial:  initial(lastName),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSuggestingEmail, err)
	}

	return strings.ToLower(buf.String()), nil
}

// initial returns the first rune of s as a string, or an empty string if s is
// empty. Indexing s[0] would panic on an empty string (the last name is
// optional) and split multi-byte names, so decode the first rune instead.
func initial(s string) string {
	if s == "" {
		return ""
	}

	r, _ := utf8.DecodeRuneInString(s)

	return string(r)
}
