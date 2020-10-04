package dijkstra

import "fmt"

type graph struct {
	ns nodes
	rs routes
}

// NewGraph instantiates a new graph structure with no routes
func NewGraph() *graph {
	return &graph{}
}

// AddRoute adds a new route to an existing graph
func (g *graph) AddRoute(from string, to string, dist int) *graph {
	g.rs = append(g.rs, route{from: from, to: to, dist: dist})

	if _, exists := g.ns.Node(from); !exists {
		g.ns = append(g.ns, node{name: from})
	}

	return g
}

// Routes returns the current routes of an existing graph
func (g graph) Routes() routes {
	return g.rs
}

// BestPath returns a path composed by the best set of routes to go from a spot to another
func (g graph) BestPath(from string, to string) (*path, error) {
	nodesByName := make(map[string]*node)

	for i := 0; i < len(g.ns); i++ {
		nodesByName[g.ns[i].name] = &g.ns[i]
	}

	destiny, exists := nodesByName[to]
	if !exists {
		return nil, fmt.Errorf("node of value %s not found", to)
	}

	origin, exists := nodesByName[from]
	if !exists {
		return nil, fmt.Errorf("node of value %s not found", from)
	}

	origin.setPath(routes{}...)

	return calculateBestPath(nodesByName, g.rs, destiny.name, origin.name)
}

// Print instantiates a print handler for an existing graph
func (g graph) Print() *graphPrintHandler {
	return &graphPrintHandler{g: g}
}

type graphPrintHandler struct {
	g graph
}

// Routes prints all current routes of an existing graph through a graphPrintHandler
func (h graphPrintHandler) Routes() {
	fmt.Println("Route # | From | To | Distance")

	for i, r := range h.g.rs {
		fmt.Println(fmt.Sprintf("%d | %s | %s | %d", i+1, r.from, r.to, r.dist))
	}
}
