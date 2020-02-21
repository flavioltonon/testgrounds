package model

type node struct {
	name string
	path path
}

func (n *node) setPath(rs ...route) *node {
	n.path.setRoutes(rs...)
	return n
}

func (n node) weight() (int, bool) {
	return n.path.Length()
}

type nodes []node

func (ns nodes) Node(name string) (*node, bool) {
	for _, n := range ns {
		if n.name != name {
			continue
		}

		return &n, true
	}

	return nil, false
}
