package sqlxselect

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/xerrors"
)

func splitPath(path string) []string {
	return strings.Split(path, ".")
}

func doubleQuote(s string) string {
	q, _ := json.Marshal(s)

	return string(q)
}

func flattenErrors(errs []error) error {
	var errStrs []string
	for i, err := range errs {
		errStrs = append(errStrs, fmt.Sprintf("error %d: %v", i, err))
	}

	return xerrors.New(strings.Join(errStrs, "\n"))
}
