package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid struct {
	width  int
	height int
	cells  [][]bool
}

func CreateGrid(width int, height int) *Grid {
	var g Grid
	g.width = width
	g.height = height
	g.cells = createCells(width, height)
	return &g
}

func createCells(width int, height int) [][]bool {
	cells := make([][]bool, height)
	for col_index, _ := range cells {
		cells[col_index] = make([]bool, width)
	}
	return cells
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
		x = g.height - 1
	} else if x == g.height {
		x = 0
	}

	if y == -1 {
		y = g.width - 1
	} else if y == g.width {
		y = 0
	}
 
	return g.cells[x][y]
}

func (g *Grid) countAliveNeighbours(x int, y int) int {
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

func (g *Grid) Tick() {
	new_cells := createCells(g.width, g.height)
	for x, row := range g.cells {
		for y, _ := range row {
			count := g.countAliveNeighbours(x, y)
			alive := g.getCell(x, y) 
			if alive && (count == 2 || count == 3) {
				new_cells[x][y] = true
			}
			if !alive && count == 3 {
				new_cells[x][y] = true
			}
		}
	}
	g.cells = new_cells
}

func (g *Grid) Print () {
	const clear_screen = "\033[2J"
	fmt.Print(clear_screen)
	for _, row := range g.cells {
		for _, cell := range row {
			if cell { 
				fmt.Print("🟩")
			} else {
				fmt.Print(" ")
			}
		}	
		fmt.Print("\n")	
	}
}

func main() {
	g := CreateGrid(80, 24)
	g.RandomlyInitialiseCells(0.1)
	for {
		g.Print()
		g.Tick()
		time.Sleep(100 * time.Millisecond)
	}
}
