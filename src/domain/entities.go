package domain

type Solution struct {
	Cars []Car
}

type Car struct {
	Capacity      int
	CurrentDemand int
	Route         Route
	Penalized     bool
}

func NewCar(carCapacity int, depotNode *Node) Car {
	return Car{
		Capacity: carCapacity,
		Route: Route{
			NodesMap: map[int]*Node{
				1: depotNode, // depot index is the first node on node map for all routes
			},
		},
	}
}

func (c *Car) SetPenalty() {
	c.Penalized = true
}

func (c *Car) RevokePenalty() {
	c.Penalized = false
}

type Route struct {
	NodesMap map[int]*Node
}

type Node struct {
	ID          int
	Demand      int
	IsDepot     bool
	CurrentEdge *Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

func NewEdge(FromNode int, ToNode int, Weight int) *Edge {
	return &Edge{
		From:   FromNode,
		To:     ToNode,
		Weight: Weight,
	}
}

type Matrix struct {
	Rows int
	Cols int
	data []float64
}

func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		Rows: rows,
		Cols: cols,
		data: make([]float64, rows*cols),
	}
}

func (m *Matrix) At(row int, col int) float64 {
	return m.data[row*col]
}

func (m *Matrix) Set(row int, col int, value float64) {
	m.data[row*col] = value
}
