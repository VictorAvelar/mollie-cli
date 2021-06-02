package displayers

import (
	"encoding/json"

	"github.com/avocatl/admiral/pkg/commander"
)

// JsonDisplayer helps you dump the whole json response
// as received from Mollie's API.
type JsonDisplayer struct {
	Data   interface{}
	Pretty bool
}

// KV is a displayable group of key value.
func (jd *JsonDisplayer) KV() []map[string]interface{} {
	var out []map[string]interface{}

	var v []byte

	var err error
	if jd.Pretty {
		v, err = json.MarshalIndent(jd.Data, "", "    ")
	} else {
		v, err = json.Marshal(jd.Data)
	}

	if err != nil {
		v = []byte(err.Error())
	}

	out = append(out, map[string]interface{}{
		"": string(v),
	})

	return out
}

// Cols returns an array of columns available for displaying.
func (jd *JsonDisplayer) Cols() []string {
	return commander.NewCols("")
}

// ColMap returns a list of columns and its description.
func (jd *JsonDisplayer) ColMap() map[string]string {
	return map[string]string{}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (jd *JsonDisplayer) NoHeaders() bool {
	return true
}
