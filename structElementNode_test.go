package sqlxselect

import (
	"reflect"
	"sort"
	"testing"
)

func prepareTree(t *testing.T) *structElementNode {
	node := &structElementNode{}

	node.addChild(splitPath("a.b.c")...)
	node.addChild(splitPath("a.b.d")...)
	node.addChild(splitPath("a.e")...)
	node.addChild(splitPath("a.e.f.g")...)

	return node
}

func TestStructElementNode(t *testing.T) {
	t.Run("findNode", func(t *testing.T) {
		node := prepareTree(t)

		if want, got := node.children["a"].children["b"], node.findNode("a", "b"); want != got {
			t.Errorf("findNode returned unexpected value: want %v, got %v", want, got)
		}

		if want, got := node, node.findNode(); want != got {
			t.Errorf("findNode returned unexpected value: want %v, got %v", want, got)
		}
	})

	t.Run("listElements", func(t *testing.T) {
		node := prepareTree(t)

		all := node.listElements()
		sort.Strings(all)

		if !reflect.DeepEqual(all, []string{"a.b.c", "a.b.d", "a.e.f.g"}) {
			t.Errorf("list for whole tree does not match: %v", all)
		}

		ab := node.findNode("a", "b").listElements()
		sort.Strings(ab)

		if !reflect.DeepEqual(ab, []string{"c", "d"}) {
			t.Errorf("list for a.b tree does not match: %v", all)
		}
	})
}
