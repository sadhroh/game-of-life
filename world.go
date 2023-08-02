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

// World represents the 2d game world where cells interact
type World [][]Cell

// New instantiates a new game world
func NewWorld(seed int64, rows, columns int) World {
	gameWorld, _ := LoadNewWorld(seed, rows, columns)
	return gameWorld
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
func (w *World) Next() {
	gameWorld := *w
	for rowIdx := 0; rowIdx < len(gameWorld); rowIdx++ {
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
		}
	}
}

// Show prints the current game state with * as live cells
func (w World) Show() {
	for _, row := range w {
		for _, column := range row {
			v := " "
			if column.IsAlive {
				v = "*"
			}
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}
}
