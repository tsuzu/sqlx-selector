package sqlxselect

import (
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
	"golang.org/x/xerrors"
)

// SqlxSelector is a generator of columns in SELECT query
type SqlxSelector struct {
	node    *structElementNode
	columns []string
	Errors  []error
}

// New generates SqlxSelector with default mapper
func New(dst interface{}) (*SqlxSelector, error) {
	return NewWithMapper(dst, reflectx.NewMapperFunc("db", strings.ToLower))
}

// NewWithMapper generates SqlxSelector with specified mapper
func NewWithMapper(dst interface{}, mapper *reflectx.Mapper) (*SqlxSelector, error) {
	m := mapper.FieldMap(reflect.ValueOf(dst))

	node := &structElementNode{}

	for path := range m {
		node.addChild(splitPath(path)...)
	}

	return &SqlxSelector{
		node: node,
	}, nil
}

// Select adds the column directly to query
func (s *SqlxSelector) Select(column string) *SqlxSelector {
	s.columns = append(s.columns, "`"+column+"`")

	return s
}

// SelectAs adds the column and 'AS' name directly to query
func (s *SqlxSelector) SelectAs(column, as string) *SqlxSelector {
	s.columns = append(s.columns, "`"+column+"` AS "+doubleQuote(as))

	return s
}

// SelectStruct adds columns specified as 'column' to query to store values
// 'limit' can specify columns to add
// ex. SelectStruct("users.*" /* table name */, "id", "name" /* columns to select */)
func (s *SqlxSelector) SelectStruct(column string, limit ...string) *SqlxSelector {
	return s.SelectStructAs(column, column, limit...)
}

// SelectStructAs adds columns specified as 'column' to query to store values specified as 'as' in struct
// 'limit' can specify columns to add
// ex. SelectStructAs("users.*" /* table name */, "user." /* 'db:""' name */, "id", "name" /* columns to select */)
func (s *SqlxSelector) SelectStructAs(column, as string, limit ...string) *SqlxSelector {
	ass := splitPath(as)

	if len(ass) != 0 && ass[len(ass)-1] == "*" {
		ass = ass[:len(ass)-1]
	}

	node := s.node.findNode(ass...)

	if node == nil {
		s.Errors = append(s.Errors, xerrors.Errorf("unknown node in %v", as))
		return s
	}

	columnPrefix := strings.TrimSuffix(column, "*")

	elms := node.listElements()

	check := true
	if len(limit) == 0 {
		check = false
		limit = elms
	}

	elmsSet := map[string]struct{}{}
	if check {
		for i := range elms {
			elmsSet[elms[i]] = struct{}{}
		}
	}

	for i := range limit {
		if check {
			_, found := elmsSet[limit[i]]

			if !found {
				s.Errors = append(s.Errors, xerrors.Errorf("unknown column: %s", limit[i]))

				continue
			}
		}

		s.SelectAs(columnPrefix+limit[i], strings.Join(append(ass, limit[i]), "."))
	}

	return s
}

func (s *SqlxSelector) String() string {
	if len(s.Errors) != 0 {
		return ""
	}
	return strings.Join(s.columns, ",")
}

// StringWithError returns columns for SELECT as string, but may return an error if something went wrong
func (s *SqlxSelector) StringWithError() (string, error) {
	if len(s.Errors) != 0 {
		return "", flattenErrors(s.Errors)
	}
	return strings.Join(s.columns, ","), nil
}
