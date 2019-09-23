package sqlxselect

type structElementNode struct {
	children map[string]*structElementNode
}

func (s *structElementNode) addChild(path ...string) {
	if len(path) == 0 {
		return
	}

	if s.children == nil {
		s.children = map[string]*structElementNode{}
	}

	v, ok := s.children[path[0]]

	if !ok {
		s.children[path[0]] = &structElementNode{}
		v = s.children[path[0]]
	}

	v.addChild(path[1:]...)
}

func (s *structElementNode) findNode(path ...string) *structElementNode {
	if len(path) == 0 {
		return s
	}

	if s.children == nil {
		return nil
	}

	v, ok := s.children[path[0]]

	if !ok {
		return nil
	}

	return v.findNode(path[1:]...)
}

func (s *structElementNode) listElements() []string {
	return s.listElementsImpl("")
}

func (s *structElementNode) listElementsImpl(ns string) []string {
	if len(s.children) == 0 {
		return []string{ns}
	}
	var res []string

	for k, v := range s.children {
		if ns != "" {
			res = append(res, v.listElementsImpl(ns+"."+k)...)
		} else {
			res = append(res, v.listElementsImpl(k)...)
		}
	}

	return res
}
