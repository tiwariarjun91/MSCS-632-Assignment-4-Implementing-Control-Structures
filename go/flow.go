
package main

// Edge represents a connection between two nodes
// in the scheduling network.
type Edge struct {
	to   int
	rev  int
	cap  int
	cost int
}

// EdgeRef keeps track of an employee's possible
// assignment to a particular day and shift.
type EdgeRef struct {
	i   int
	d   int
	s   int
	u   int
	idx int
}

// addEdge connects two nodes in the scheduling graph.
//
// A forward edge represents an available assignment.
// A reverse edge allows the algorithm to undo or
// rearrange an earlier assignment when necessary.
func addEdge(
	graph *[][]Edge,
	from int,
	to int,
	capacity int,
	cost int,
) int {

	forwardIndex := len((*graph)[from])
	reverseIndex := len((*graph)[to])

	forwardEdge := Edge{
		to:   to,
		rev:  reverseIndex,
		cap:  capacity,
		cost: cost,
	}

	reverseEdge := Edge{
		to:   from,
		rev:  forwardIndex,
		cap:  0,
		cost: -cost,
	}

	(*graph)[from] = append(
		(*graph)[from],
		forwardEdge,
	)

	(*graph)[to] = append(
		(*graph)[to],
		reverseEdge,
	)

	return forwardIndex
}
