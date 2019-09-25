package sqlxselect

import (
	"encoding/json"
	"strings"
)

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

func doubleQuote(s string) string {
	q, _ := json.Marshal(s)

	return string(q)
}
