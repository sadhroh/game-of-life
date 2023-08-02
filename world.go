package main

import "fmt"

type Cell struct {
	IsAlive  bool
	WasAlive bool
}

func NewCell(isAlive bool) Cell {
	return Cell{
		IsAlive: isAlive,
	}
}

type World [][]Cell

func NewWorld() World {
	return World{}
}

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
