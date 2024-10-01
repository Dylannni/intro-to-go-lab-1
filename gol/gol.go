package main

func countLiveNeighbors(world [][]byte, row, col int) int {
	rows := len(world)
	cols := len(world[0])
	neighbors := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top-left, Top, Top-right
		{0, -1}, {0, 1}, // Left,            Right
		{1, -1}, {1, 0}, {1, 1}, // Bottom-left, Bottom, Bottom-right
	}

	liveNeighbors := 0
	for _, n := range neighbors {
		newRow := (row + n[0] + rows) % rows
		newCol := (col + n[1] + cols) % cols
		if world[newRow][newCol] == 255 {
			liveNeighbors++
		}
	}
	return liveNeighbors
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	rows := len(world)
	cols := len(world[0])

	// Create a new world for the next state
	newWorld := make([][]byte, rows)
	for i := range newWorld {
		newWorld[i] = make([]byte, cols)
	}

	// Iterate over each cell in the world
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Count the live neighbors
			liveNeighbors := countLiveNeighbors(world, row, col)

			// Apply the Game of Life rules
			if world[row][col] == 255 {
				// Cell is alive
				if liveNeighbors < 2 || liveNeighbors > 3 {
					newWorld[row][col] = 0 // Cell dies
				} else {
					newWorld[row][col] = 255 // Cell stays alive
				}
			} else {
				// Cell is dead
				if liveNeighbors == 3 {
					newWorld[row][col] = 255 // Cell becomes alive
				} else {
					newWorld[row][col] = 0 // Cell stays dead
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var aliveCells []cell

	// Iterate over every cell in the world
	for row := 0; row < len(world); row++ {
		for col := 0; col < len(world[0]); col++ {
			// If the cell is alive (value is 255), add it to the list
			if world[row][col] == 255 {
				aliveCells = append(aliveCells, cell{x: col, y: row})
			}
		}
	}

	return aliveCells
}
