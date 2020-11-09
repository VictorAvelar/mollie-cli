package runners

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNopRunner(t *testing.T) {
	cmd := &cobra.Command{
		Run: NopRunner,
	}

	emptyBytes := []byte{}

	b := bytes.NewBuffer(emptyBytes)

	cmd.SetOut(b)
	assert.Nil(t, cmd.Execute())

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
	}

	assert.Empty(t, out)
}

func TestArgPrinterRunner(t *testing.T) {
	cases := []struct {
		name   string
		expect string
		input  []string
	}{
		{
			name:   "hello world as args",
			expect: "hello world",
			input:  []string{"hello", "world"},
		},
		{
			name:   "empty or no args",
			expect: "",
			input:  []string{},
		},
		{
			name:   "some long text",
			expect: "Hugo is a very fast static site generator",
			input:  strings.Split("Hugo is a very fast static site generator", " "),
		},
	}

	cmd := &cobra.Command{
		Run: ArgPrinterRunner,
	}

	for _, c := range cases {
		b := bytes.NewBufferString("")
		cmd.SetOutput(b)
		t.Run(c.name, func(t *testing.T) {
			cmd.SetArgs(c.input)
			cmd.Execute()
			out, err := ioutil.ReadAll(b)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, c.expect, string(out))
		})
	}
}
