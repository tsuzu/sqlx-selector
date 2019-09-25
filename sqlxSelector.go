package sqlxselect

import (
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
	"golang.org/x/xerrors"
)

type SqlxSelector struct {
	node    *structElementNode
	columns []string
	err     error
}

func New(dst interface{}) (*SqlxSelector, error) {
	return NewWithMapper(dst, reflectx.NewMapperFunc("db", strings.ToLower))
}

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

func (s *SqlxSelector) Select(column string) *SqlxSelector {
	s.columns = append(s.columns, "`"+column+"`")

	return s
}

func (s *SqlxSelector) SelectAs(column, as string) *SqlxSelector {
	s.columns = append(s.columns, "`"+column+"` AS "+doubleQuote(as))

	return s
}

func (s *SqlxSelector) SelectStruct(column string) *SqlxSelector {
	return s.SelectStructAs(column, column)
}

func (s *SqlxSelector) SelectStructAs(column, as string) *SqlxSelector {
	ass := splitPath(as)

	if len(ass) != 0 && ass[len(ass)-1] == "*" {
		ass = ass[:len(ass)-1]
	}

	node := s.node.findNode(ass...)

	if node == nil {
		s.err = xerrors.Errorf("unknown node in %v", as)
		return s
	}

	columnPrefix := strings.TrimSuffix(column, "*")

	elms := node.listElements()

	for i := range elms {
		s.SelectAs(columnPrefix+elms[i], strings.Join(append(ass, elms[i]), "."))
	}

	return s
}

func (s *SqlxSelector) String() string {
	if s.err != nil {
		return ""
	}
	return strings.Join(s.columns, ",")
}

func (s *SqlxSelector) StringWithError() (string, error) {
	if s.err != nil {
		return "", s.err
	}
	return strings.Join(s.columns, ","), nil
}
