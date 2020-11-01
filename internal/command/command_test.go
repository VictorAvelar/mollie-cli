package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	config := Config{
		Namespace: "test",
		Example:   "test <name>",
		Hidden:    false,
	}

	c := Builder(nil, config, []string{"test1", "col2"})

	assert.Equal(t, c.Command.Name(), config.Namespace)
	assert.False(t, c.Hidden)
	assert.Equal(t, c.Example, config.Example)
}

func TestBuilder_WithParent(t *testing.T) {
	cp := Config{
		Namespace: "parent",
		Example:   "parent <name>",
	}

	p := Builder(nil, cp, []string{})

	child := Builder(p, Config{Namespace: "child"}, []string{})

	assert.Equal(t, p.Command, child.Parent())
	assert.Nil(t, p.Parent())
	assert.ElementsMatch(t, []*Command{child}, p.GetSubCommands())
}

func TestGetCols(t *testing.T) {
	cols := []string{"test"}
	c := Builder(nil, Config{Namespace: "test"}, cols)

	assert.Equal(t, cols, c.Cols)
}

func TestAddFlag_String(t *testing.T) {
	cmd := Builder(nil, Config{Namespace: "test"}, []string{})

	AddFlag(cmd, FlagConfig{
		FlagType:  StringFlag,
		Name:      "test-flag",
		Shorthand: "t",
		Default:   "default",
	})

	flag, err := cmd.Flags().GetString("test-flag")
	assert.Nil(t, err)
	assert.Equal(t, "default", flag)
	assert.IsType(t, "", flag)
}

func TestAddFlag_Int(t *testing.T) {
	cmd := Builder(nil, Config{Namespace: "test"}, []string{})

	AddFlag(cmd, FlagConfig{
		FlagType:  IntFlag,
		Name:      "test-flag",
		Shorthand: "t",
		Default:   10,
	})

	flag, err := cmd.Flags().GetInt("test-flag")
	assert.Nil(t, err)
	assert.Equal(t, 10, flag)
	assert.IsType(t, 0, flag)
}

func TestAddFlag_Int64(t *testing.T) {
	cmd := Builder(nil, Config{Namespace: "test"}, []string{})

	AddFlag(cmd, FlagConfig{
		FlagType:  Int64Flag,
		Name:      "test-flag",
		Shorthand: "t",
		Default:   int64(10),
	})

	flag, err := cmd.Flags().GetInt64("test-flag")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), flag)
	assert.IsType(t, int64(10), flag)
}

func TestAddFlag_Bool(t *testing.T) {
	cmd := Builder(nil, Config{Namespace: "test"}, []string{})

	AddFlag(cmd, FlagConfig{
		FlagType:  BoolFlag,
		Name:      "test-flag",
		Shorthand: "t",
		Default:   true,
	})

	flag, err := cmd.Flags().GetBool("test-flag")
	assert.Nil(t, err)
	assert.Equal(t, true, flag)
	assert.IsType(t, false, flag)
}

func TestAddFlag_Float64(t *testing.T) {
	cmd := Builder(nil, Config{Namespace: "test"}, []string{})

	AddFlag(cmd, FlagConfig{
		FlagType:  Float64Flag,
		Name:      "test-flag",
		Shorthand: "t",
		Default:   float64(10.15),
	})

	flag, err := cmd.Flags().GetFloat64("test-flag")
	assert.Nil(t, err)
	assert.Equal(t, float64(10.15), flag)
	assert.IsType(t, float64(0), flag)
}

func TestAddFlag_Persistent(t *testing.T) {
	parent := Builder(nil, Config{Namespace: "test"}, []string{})

	AddFlag(parent, FlagConfig{
		FlagType:   BoolFlag,
		Name:       "test-flag",
		Shorthand:  "t",
		Default:    true,
		Persistent: true,
	})

	child := Builder(parent, Config{Namespace: "child-test"}, []string{})

	flag, err := child.Parent().PersistentFlags().GetBool("test-flag")
	assert.Nil(t, err)
	assert.Equal(t, true, flag)
	assert.IsType(t, false, flag)
}
