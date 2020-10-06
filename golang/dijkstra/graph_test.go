package dijkstra_test

import (
	"testgrounds/dijkstra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	g := dijkstra.NewGraph()

	g.AddRoute("A", "B", 2)
	g.AddRoute("A", "C", 3)
	g.AddRoute("B", "C", 2)
	g.AddRoute("B", "D", 4)
	g.AddRoute("C", "D", 6)
	g.AddRoute("D", "A", 8)

	assert.Equal(t, 6, len(g.Routes()))

	p, err := g.BestPath("A", "C")
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
