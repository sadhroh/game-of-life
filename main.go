package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// command line flags
var (
	gameFile         = flag.String("file", "docs/world.txt", "current save state of the world")
	gameNew          = flag.Bool("new", true, "generate new game world")
	gameSeed         = flag.Int64("seed", time.Now().UnixNano(), "starting seed")
	gameWorldRows    = flag.Int("rows", 40, "number of rows")
	gameWorldColumns = flag.Int("columns", 40, "number of columns")
)

var universe World

func init() {
	flag.Parse()
	var err error
	// initialise new world
	universe, err = NewWorld(&Spec{
		New:      *gameNew,
		GameFile: *gameFile,
		Seed:     *gameSeed,
		Rows:     *gameWorldRows,
		Columns:  *gameWorldColumns,
	})
	if err != nil {
		log.Fatalln("could not start game", err)
	}
}

func Evolve() {
	generation := 0
	
	for {
		fmt.Println("Generation:", generation)
		universe.Show()
		universe.Next()
		time.Sleep(300 * time.Millisecond)
		generation++
	}
}

func main() {
	Evolve()
}
