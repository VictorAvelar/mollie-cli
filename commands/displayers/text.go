package displayers

import (
	"fmt"
	"strings"

	"github.com/avocatl/admiral/pkg/commander"
	"github.com/avocatl/admiral/pkg/display"
)

// TextDisplayer for printing messaged during as
// admiral displayable printers.
type TextDisplayer struct {
	Divider string
	Text    string
}

// KV is a displayable group of key value.
func (td *TextDisplayer) KV() []map[string]interface{} {
	var out []map[string]interface{}

	br := strings.Repeat(td.Divider, 50)

	text := fmt.Sprintf("\n%s\n%s\n%s\n",
		br,
		strings.TrimSuffix(td.Text, "\n"),
		br,
	)

	out = append(out, map[string]interface{}{
		"": text,
	})
	return out
}

// Cols returns an array of columns available for displaying.
func (td *TextDisplayer) Cols() []string {
	return commander.NewCols("")
}

// ColMap returns a list of columns and its description.
func (td *TextDisplayer) ColMap() map[string]string {
	return map[string]string{}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (td *TextDisplayer) NoHeaders() bool {
	return true
}

// NewSimpleTextDisplayer programatically builds a simple text
// displayer.
func NewSimpleTextDisplayer(divider, text string) display.Displayable {
	return &TextDisplayer{
		Divider: divider,
		Text:    text,
	}
}
