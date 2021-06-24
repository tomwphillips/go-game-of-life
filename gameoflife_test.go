package main

import "testing"

func TestCreateGrid(t *testing.T) {
	width := 3
	height := 5
	grid := CreateGrid(width, height)

	if len(grid.cells) != height {
		t.Errorf(`len(grid.cells) = %v, want %v`, len(grid.cells), height)
	}

	for i, row := range grid.cells {
		if len(row) != width {
			t.Errorf(`len(grid.cells[%v]) = %v, want %v`, i, len(row), width)
		}
	}
}

func countAliveCells(g *Grid) int {
	alive_count := 0
	for _, row := range g.cells {
		for _, value := range row {
			if value {
				alive_count += 1
			}
		}
	}
	return alive_count
}

func TestRandomlyInitialiseCells(t *testing.T) {
	width := 10
	height := 10
	g := CreateGrid(width, height)

	probability := 0.0
	g.RandomlyInitialiseCells(probability)
	if alive_count := countAliveCells(g); alive_count != 0 {
		t.Errorf(`alive count = %v, want 0 for probability = %v`, alive_count, probability)
	}

	probability = 1.0
	g.RandomlyInitialiseCells(probability)
	if alive_count := countAliveCells(g); alive_count != width*height {
		t.Errorf(`alive count = %v, want %v for probability = %v`, alive_count, width*height, probability)
	}
}

func TestGetCell(t *testing.T) {
	g := CreateGrid(2, 2)
	g.cells[0][0] = true
	g.cells[1][1] = true

	tests := []struct{
		x int
		y int
		want bool
	}{
		{0, 0, true},
		{0, 1, false},
		{1, 0, false},
		{1, 1, true},
		{-1, 0, false},
		{2, 0, true},
		{0, -1, false},
		{0, 2, true},
	}

	for _, test := range tests {
		if got := g.getCell(test.x, test.y); got != test.want {
			t.Errorf(`g.Get(%v, %v) = %v, want %v`, test.x, test.y, got, test.want)
		}
	}
}

func TestCountAliveNeighbours(t *testing.T) {
	g := CreateGrid(2, 2)
	g.cells[0][0] = true
	g.cells[1][1] = true

	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			if got := g.CountAliveNeighbours(x, y); got != 4 {
				t.Errorf(`g.CountNeighbours(%v, %v) = %v, want 4`, x, y, got)
			}
		}
	}
}

func TestTick(t *testing.T) {
	g := CreateGrid(5, 5)
	g.Tick()
}