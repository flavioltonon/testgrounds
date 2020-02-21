package model

import (
	"fmt"

	"github.com/pkg/errors"
)

type path struct {
	rs    routes
	isSet bool
}

func (p *path) setRoutes(rs ...route) *path {
	p.rs = rs
	p.isSet = true
	return p
}

// Length returns the total distance of an existing path
func (p path) Length() (int, bool) {
	if len(p.rs) == 0 && !p.isSet {
		return 0, false
	}

	var dist int

	for _, r := range p.rs {
		dist += r.dist
	}

	return dist, true
}

func calculateBestPath(nodesByName map[string]*node, rs routes, destiny string, origins ...string) (*path, error) {
	nextNodes := make([]string, 0)

	remainingRoutes := make(routes, len(rs))

	copy(remainingRoutes, rs)

	for i := 0; i < len(origins); i++ {
		from, exists := nodesByName[origins[i]]
		if !exists {
			continue
		}

		if from.name == destiny {
			return &from.path, nil
		}

		for _, r := range rs {
			// find current node
			if r.from != from.name {
				continue
			}

			to, exists := nodesByName[r.to]
			if !exists {
				continue
			}

			tw, isSet := to.path.Length()
			fw, _ := from.path.Length()

			if !isSet || tw > fw+r.dist {
				to.setPath(append(from.path.rs, r)...)
			}

			remainingRoutes = remainingRoutes.pop(from.name, to.name)
			nextNodes = append(nextNodes, to.name)
		}
	}
	if len(nextNodes) == 0 {
		return nil, errors.Wrap(errors.New("missing routes and/or nodes"), "unable to calculate best path")
	}

	return calculateBestPath(nodesByName, remainingRoutes, destiny, nextNodes...)
}

// Print instantiates a print handler for an existing path
func (p path) Print() *pathPrintHandler {
	return &pathPrintHandler{p: p}
}

type pathPrintHandler struct {
	p path
}

// Routes prints all current routes of an existing path through a pathPrintHandler
func (h pathPrintHandler) Routes() {
	fmt.Println("Route # | From | To | Distance")

	for i, r := range h.p.rs {
		fmt.Println(fmt.Sprintf("%d | %s | %s | %d", i+1, r.from, r.to, r.dist))
	}
}
