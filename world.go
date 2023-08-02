package main

import "fmt"

// Cell represents the single entity
type Cell struct {
	// IsAlive stores the current state of the cell
	IsAlive bool
	// WasAlive stores cell state, before it's latest state change
	WasAlive bool
}

// NewCell creates a new cell with the state as specified in the argument
func NewCell(isAlive bool) Cell {
	return Cell{
		IsAlive: isAlive,
	}
}

// Spec is the game universe specification
type Spec struct {
	// Rows specifies the number of rows in the 2d game world
	Rows int
	// Columns specifies the number of rows in the 2d game world
	Columns int
	// Seed can be given as the starting random value to generate the new world
	Seed int64
	// GameFile contains the path to the save state of the game world
	GameFile string
}

// World represents the 2d game world where cells interact
type World [][]Cell

// Universe represents an instance of the 2d game world
type Universe struct {
	rows      int
	cols      int
	liveCells int
	world     [][]Cell
}

// New instantiates a new game world with the specification as provided in the argument
func NewUniverse(cfg *Spec) (Universe, error) {
	// load game universe from save file
	if cfg.GameFile != "" {
		universe, err := LoadWorldFromFile(cfg.GameFile)
		if err != nil {
			return Universe{}, fmt.Errorf("unable to load the game world: %+w", err)
		}
		return universe, nil
	}

	// generate new game univers
	universe, _ := LoadNewWorld(cfg.Seed, cfg.Rows, cfg.Columns)
	return universe, nil
}

// Next generates the next state of the game world in place.
// The cells interact as follows:
//	 * a living cell with less than 2 live neighbours dies --> underpopulation
//	 * a living cell with 2 or 3 live neighbours           --> continues to live
//	 * a living cell with more than 3 live neighbours dies --> overcrowding
//   * a dead cell with 3 live neighbours, lives           --> reproduction
//
// A cell has 8 neighbours.
// If Cell(x,y) represents the co-ordinates of the cell, the neighbour map relative to this cell is:
//		(x-1, y-1)  (x-1, y)  (x-1, y+1)
//		(x, y-1)    [ Cell ]   (x, y+1)
//		(x+1, y-1)  (x+1, y)  (x+1, y+1)
func (w *Universe) Next() {
	livingCells := 0
	gameWorld := (*w).world
	for rowIdx := 0; rowIdx < w.rows; rowIdx++ {
		for colIdx := 0; colIdx < len(gameWorld[rowIdx]); colIdx++ {
			isAlive := gameWorld[rowIdx][colIdx].IsAlive

			aliveNeighbours := 0

			// check above row neighbours
			if rowIdx-1 >= 0 && colIdx-1 >= 0 &&
				gameWorld[rowIdx-1][colIdx-1].WasAlive {
				aliveNeighbours++
			}
			if rowIdx-1 >= 0 &&
				gameWorld[rowIdx-1][colIdx].WasAlive {
				aliveNeighbours++
			}
			if rowIdx-1 >= 0 && colIdx+1 < len(gameWorld[rowIdx]) &&
				gameWorld[rowIdx-1][colIdx+1].WasAlive {
				aliveNeighbours++
			}

			// check same row neighbours
			if colIdx-1 >= 0 &&
				gameWorld[rowIdx][colIdx-1].WasAlive {
				aliveNeighbours++
			}
			if colIdx+1 < len(gameWorld[rowIdx]) &&
				gameWorld[rowIdx][colIdx+1].IsAlive {
				aliveNeighbours++
			}

			// check below row neighbours
			if rowIdx+1 < len(gameWorld) && colIdx-1 >= 0 &&
				gameWorld[rowIdx+1][colIdx-1].IsAlive {
				aliveNeighbours++
			}
			if rowIdx+1 < len(gameWorld) &&
				gameWorld[rowIdx+1][colIdx].IsAlive {
				aliveNeighbours++
			}
			if rowIdx+1 < len(gameWorld) && colIdx+1 < len(gameWorld[rowIdx]) &&
				gameWorld[rowIdx+1][colIdx+1].IsAlive {
				aliveNeighbours++
			}

			willBeAlive := false                     // default dead when due to overcrowding or currently dead
			if (!isAlive && aliveNeighbours == 3) || // reproduce
				(isAlive && (aliveNeighbours == 2 || aliveNeighbours == 3)) { // continue to live
				willBeAlive = true
			}

			// update cell state
			gameWorld[rowIdx][colIdx].IsAlive = willBeAlive
			gameWorld[rowIdx][colIdx].WasAlive = isAlive

			if willBeAlive {
				livingCells++
			}
		}
	}
	w.liveCells = livingCells
}

// Show prints the current game state with * as live cells
func (w Universe) Show() {
	// to pretty print top & bottom borders
	borderPrinter := func(cols int) {
		for i := 0; i <= cols; i++ {
			fmt.Print("==")
		}
		fmt.Println()
	}

	borderPrinter(w.cols)       // for top border
	defer borderPrinter(w.cols) // for bottom border

	for _, row := range w.world {
		fmt.Print("|")
		for _, column := range row {
			v := " "
			if column.IsAlive {
				v = "*"
			}
			fmt.Printf("%s ", v)
		}
		fmt.Println("|")
	}
}

// Dead returns if all cells are dead in the universe
func (w Universe) Dead() bool {
	return w.liveCells == 0
}
