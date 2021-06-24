package main

import "math/rand"

type Grid struct {
	width  int
	height int
	cells  [][]bool
}

func CreateGrid(width int, height int) *Grid {
	var g Grid
	g.height = height
	g.width = width
	g.cells = make([][]bool, height)
	for col_index, _ := range g.cells {
		g.cells[col_index] = make([]bool, width)
	}
	return &g
}

func (g *Grid) RandomlyInitialiseCells(probability float64) *Grid {
	for x, row := range g.cells {
		for y, _ := range row {
			g.cells[x][y] = rand.Float64() < probability
		}
	}
	return g
}

func (g *Grid) getCell(x int, y int) bool {
	if x == -1 {
		x = g.width - 1
	} else if x == g.width {
		x = 0
	}

	if y == -1 {
		y = g.height - 1
	} else if y == g.height {
		y = 0
	}
 
	return g.cells[x][y]
}

func (g *Grid) CountAliveNeighbours(x int, y int) int {
	var count int
	offsets := []int{-1, 0, 1}
	for _, offset_x := range offsets {
		for _, offset_y := range offsets {
			if offset_x == 0 && offset_y == 0 {
				continue
			}
			if g.getCell(x + offset_x, y + offset_y) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) Tick() *Grid{
	return &Grid{}
}

func main() {
	g := CreateGrid(3, 4)
	g.RandomlyInitialiseCells(0.5)
	for i := 0; i < 100; i++ {
		g.Tick()
	}
}
