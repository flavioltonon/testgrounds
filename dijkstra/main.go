package main

import (
	"flavioltonon/testgrounds/dijkstra/model"
)

func main() {
	g := model.NewGraph()

	g.AddRoute("A", "B", 2)
	g.AddRoute("A", "C", 3)
	g.AddRoute("B", "C", 2)
	g.AddRoute("B", "D", 4)
	g.AddRoute("C", "D", 6)
	g.AddRoute("D", "A", 8)

	g.Print().Routes()

	p, err := g.BestPath("A", "A")
	if err != nil {
		panic(err)
	}

	p.Print().Routes()
}
