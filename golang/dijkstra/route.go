package dijkstra

type route struct {
	from string
	to   string
	dist int
}

type routes []route

func (r routes) pop(from, to string) routes {
	for i := 0; i < len(r); i++ {
		if r[i].from != from {
			continue
		}

		if r[i].to != to {
			continue
		}

		return append(r[:i], r[i+1:]...)
	}

	return r
}
