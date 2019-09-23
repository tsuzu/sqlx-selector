package sqlxselect

import "strings"

func splitPath(path string) []string {
	return strings.Split(path, ".")
}
