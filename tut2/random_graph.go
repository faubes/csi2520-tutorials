package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type graph struct {
	vertices []*vertex
	edges    []*edge
	order    int // |V|
	size     int // |E|
	lock     sync.RWMutex
}

type vertex struct {
	neighbours []*vertex
	value      int
	visited    bool
	lock       sync.RWMutex
}

type edge struct {
	endpoints []*vertex
	lock      sync.RWMutex
}

func (g *graph) getSize() int {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.size
}

func (g *graph) getOrder() int {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.order
}

func (g *graph) getVertex(i int) *vertex {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.vertices[i]
}

func (g *graph) hasEdge(v1 *vertex, v2 *vertex) bool {
	g.lock.RLock()
	defer g.lock.RUnlock()
	for _, e := range g.edges {
		if (e.endpoints[0] == v1 &&
			e.endpoints[1] == v2) ||
			(e.endpoints[0] == v2 &&
				e.endpoints[1] == v1) {
			return true
		}
	}
	return false
}

func (g *graph) NewVertex() *vertex {
	g.lock.Lock()
	defer g.lock.Unlock()
	v := &vertex{}
	v.lock.Lock()
	v.value = g.order
	v.lock.Unlock()
	g.vertices = append(g.vertices, v)
	g.order++
	return v
}

func (g *graph) AddEdge(v1 *vertex, v2 *vertex) {
	g.lock.Lock()
	defer g.lock.Unlock()
	new_edge := &edge{}
	new_edge.lock.Lock()
	new_edge.endpoints = []*vertex{v1, v2}
	g.edges = append(g.edges, new_edge)
	new_edge.lock.Unlock()
	v1.lock.Lock()
	v1.neighbours = append(v1.neighbours, v2)
	v1.lock.Unlock()
	v2.lock.Lock()
	v2.neighbours = append(v2.neighbours, v1)
	v2.lock.Unlock()
	g.size++
}

func (g *graph) random_walker(v *vertex, p float32,
	wg *sync.WaitGroup, maxdepth int) {
	defer wg.Done()
	if maxdepth <= 0 {
		return // stop after reaching max depth
	}
	// x random uniform in [0, 1)
	x := rand.Float32()
	// fmt.Println("Random roll is ", x)
	var u *vertex // next vertex in walk
	switch {
	// create a new node with probability p
	case x <= p:
		u = g.NewVertex()
		g.AddEdge(u, v)
	default:
		// add an edge to a random vertex
		// (if it doesn't already exist)
		// and walk there
		n := g.getOrder()
		i := rand.Intn(n)
		u = g.getVertex(i)
		if u == v {
			// don't make loops
			break
		}
		if !g.hasEdge(u, v) {
			g.AddEdge(u, v)
		}
	}

	go func() {
		wg.Add(1)
		g.random_walker(u, p, wg, maxdepth-1)
	}()
}

func (g *graph) String() string {
	str := fmt.Sprintf("Order: %d\n", g.getOrder())
	str += fmt.Sprintf("Size: %d\n", g.getSize())
	str += "Vertices: "
	for _, v := range g.vertices {
		str += strconv.Itoa(v.value)
		str += " "
	}
	str += "\n"
	str += "Edges: "
	for _, e := range g.edges {
		str += "{"
		str += strconv.Itoa(e.endpoints[0].value)
		str += ","
		str += strconv.Itoa(e.endpoints[1].value)
		str += "}"
	}
	str += "\n"
	return str
}

func (v *vertex) DFS() {
	if v.visited == true {
		return
	}
	fmt.Print(v.value)
	fmt.Print(": ")
	for _, u := range v.neighbours {
		fmt.Printf("%d ", u.value)
	}
	fmt.Println()
	v.visited = true
	for _, u := range v.neighbours {
		u.DFS()
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	g := graph{}
	v := g.NewVertex()
	fmt.Println("Generating random graph")
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		//fmt.Println("Starting walker ", i)
		go func() {
			g.random_walker(v, 0.65, &wg, 5)
		}()
	}
	// wait for all walkers to finish
	wg.Wait()
	// run DFS
	fmt.Println("Depth first search from root gives:")
	v.DFS()
	fmt.Println(g.String())
}
