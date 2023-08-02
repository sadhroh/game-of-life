package main

import (
	"math/rand"
)

// LoadNewWorld generates a new 2d game world with the seed
// additionally taking in the number of rows & columns of the 2d array world.
func LoadNewWorld(seed int64, gameWorldRows, gameWorldColumns int) (World, error) {
	rand.Seed(seed)

	gameWorld := make(World, gameWorldRows)

	for row := 0; row < gameWorldRows; row++ {
		gameRow := make([]Cell, gameWorldColumns)
		for column := 0; column < gameWorldColumns; column++ {
			gameRow[column] = NewCell(rand.Intn(2) > 0)
		}
		gameWorld[row] = gameRow
	}
	return gameWorld, nil
}
