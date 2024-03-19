package main

import "fmt"

type Vertex struct {
	value int
	adj   []*Vertex
}
type Graph struct {
	vers []*Vertex
}

func (g *Graph) AddVertex(d int) *Vertex {
	for _, v := range g.vers {
		if v.value == d {
			return v
		}
	}
	newvertex := &Vertex{value: d}
	g.vers = append(g.vers, newvertex)
	return newvertex
}

func (g *Graph) AddEdge(from, to *Vertex) {
	for _, edge := range from.adj {
		if edge == to {
			return
		}
	}
	from.adj = append(from.adj, to)
}

func (g *Graph) BFS(startVertex *Vertex) {
	visited := make(map[int]bool)
	queue := []*Vertex{startVertex}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if !visited[current.value] {
			visited[current.value] = true
			fmt.Printf("Visited: %d\n", current.value)
			for _, edge := range current.adj {
				if !visited[edge.value] {
					queue = append(queue, edge)
				}
			}
		}
	}
}

func (g *Graph) DFS(startVertex *Vertex) {
	visited := make(map[int]bool)
	stack := []*Vertex{startVertex}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[current.value] {
			visited[current.value] = true
			fmt.Printf("Visited: %d\n", current.value)
			for _, edge := range current.adj {
				if !visited[edge.value] {
					stack = append(stack, edge)
				}
			}
		}
	}
}

func main() {
	graph := Graph{}
	vertex1 := graph.AddVertex(1)
	vertex2 := graph.AddVertex(2)
	vertex3 := graph.AddVertex(3)
	vertex4 := graph.AddVertex(4)
	vertex5 := graph.AddVertex(5)

	graph.AddEdge(vertex1, vertex2)
	graph.AddEdge(vertex1, vertex3)
	graph.AddEdge(vertex2, vertex3)
	graph.AddEdge(vertex1, vertex5)
	graph.AddEdge(vertex4, vertex5)
	graph.AddEdge(vertex3, vertex1)

	//perform DFS starting from vertex
	fmt.Println("BFS using stack:")
	graph.BFS(vertex1)
	fmt.Println()

	//perform BFS starting from vertex
	fmt.Println("BFS using queue:")
	graph.DFS(vertex1)

}
