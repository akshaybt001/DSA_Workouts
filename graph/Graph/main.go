package main

import "fmt"

type Vertex struct {
	value    int
	adjecent []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) AddVertex(v int) *Vertex {
	for _, val := range g.vertices {
		if val.value == v {
			return val
		}
	}
	newNode := &Vertex{value: v}
	g.vertices = append(g.vertices, newNode)
	return newNode
}

func (g *Graph) AddEdges(from, to *Vertex) {
	for _, edge := range from.adjecent {
		if edge == to {
			return
		}
	}
	from.adjecent = append(from.adjecent, to)
}
func (g *Graph) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("%d =>", vertex.value)
		for _, edge := range vertex.adjecent {
			fmt.Printf(" %d ", edge.value)
		}
		fmt.Println()
	}
}

func main() {
	graph := &Graph{}
	ver1 := graph.AddVertex(1)
	ver2 := graph.AddVertex(2)
	ver3 := graph.AddVertex(3)
	ver4 := graph.AddVertex(4)

	graph.AddEdges(ver1, ver2)
	graph.AddEdges(ver1, ver3)
	graph.AddEdges(ver2, ver2)
	graph.AddEdges(ver3, ver4)
	graph.AddEdges(ver4, ver3)

	graph.Print()
}
