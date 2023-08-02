package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// LoadNewWorld generates a new 2d game world with the seed
// additionally taking in the number of rows & columns of the 2d array world.
func LoadNewWorld(seed int64, gameWorldRows, gameWorldColumns int) (Universe, error) {
	rand.Seed(seed)

	gameWorld := make(World, gameWorldRows)

	for row := 0; row < gameWorldRows; row++ {
		gameRow := make([]Cell, gameWorldColumns)
		for column := 0; column < gameWorldColumns; column++ {
			gameRow[column] = NewCell(rand.Intn(2) > 0)
		}
		gameWorld[row] = gameRow
	}
	return Universe{
		world: gameWorld,
		rows:  gameWorldRows,
		cols:  gameWorldColumns,
	}, nil
}

// LoadWorldFromFile parses & loads the game world into memory.
// It assumes the game world to be in a 2d integer array format
// with every successful number treated as living cell(1) & the other
// index elements treated as dead cells(0), separated by spaces.
func LoadWorldFromFile(filePath string) (Universe, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return Universe{}, fmt.Errorf("could not read file<%s>: %w", filePath, err)
	}
	defer f.Close()

	gameWorld := make(World, 0)
	cols := 0

	// buffer the read operation in case of very large files
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		cellsInRow := strings.Split(strings.TrimSpace(sc.Text()), " ")
		gameRow := make([]Cell, len(cellsInRow))
		if len(cellsInRow) > cols {
			cols = len(cellsInRow)
		}
		// for every cell, parse numbers to set as alive or dead cell
		for cellIdx, cell := range cellsInRow {
			val, _ := strconv.Atoi(cell)
			if val > 0 { // treat errors as dead cells
				val = 1
			}
			gameRow[cellIdx] = NewCell(val > 0)
		}
		gameWorld = append(gameWorld, gameRow)
	}
	if err = sc.Err(); err != nil {
		return Universe{}, fmt.Errorf("data read failed: %w", err)
	}
	return Universe{
		world: gameWorld,
		rows:  len(gameWorld),
		cols:  cols,
	}, nil
}
