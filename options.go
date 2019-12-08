package sqlxselect

import "strings"

// ColumnEscaper is a helper to escape column names
type ColumnEscaper func(s string) string

var (
	// Backquote for MySQL
	Backquote ColumnEscaper = func(s string) string {
		spl := strings.Split(s, ".")

		for i := range spl {
			spl[i] = "`" + spl[i] + "`"
		}
		return strings.Join(spl, ".")
	}

	// Doublequote for SQLite and PostgreSQL
	Doublequote ColumnEscaper = func(s string) string {
		spl := strings.Split(s, ".")

		for i := range spl {
			spl[i] = doubleQuote(spl[i])
		}
		return strings.Join(spl, ".")
	}
)

// DefaultColumnEscaper is used in New/NewWithMapper
var DefaultColumnEscaper ColumnEscaper

func init() {
	DefaultColumnEscaper = Backquote
}
