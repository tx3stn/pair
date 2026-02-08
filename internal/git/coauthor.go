package git

import "fmt"

type CoAuthor struct {
	Name  string
	Email string
}

func (c CoAuthor) Format() string {
	return fmt.Sprintf("Co-authored-by: %s <%s>", c.Name, c.Email)
}
