package sqlxselect

import (
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
)

type SqlxSelector struct {
	node *structElementNode
}

func New(arg interface{}) (*SqlxSelector, error) {
	return NewWithMapper(arg, reflectx.NewMapperFunc("db", strings.ToLower))
}

func NewWithMapper(arg interface{}, mapper *reflectx.Mapper) (*SqlxSelector, error) {
	m := mapper.FieldMap(reflect.ValueOf(arg))

	node := &structElementNode{}

	for path := range m {
		node.addChild(splitPath(path)...)
	}

	return &SqlxSelector{
		node: node,
	}, nil
}

func Select(arg interface{}, selectColumns ...string) {

}
