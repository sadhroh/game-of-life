package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// command line flags
var (
	gameFile         = flag.String("file", "", "current save state of the world")
	gameSeed         = flag.Int64("seed", time.Now().UnixNano(), "starting seed")
	gameWorldRows    = flag.Int("rows", 40, "number of rows")
	gameWorldColumns = flag.Int("columns", 40, "number of columns")
)

var universe Universe

func init() {
	flag.Parse()
	var err error
	// initialise new universe
	universe, err = NewUniverse(&Spec{
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

	if *gameFile != ""{
		*gameSeed = 0
	}
	// infinitely iterate & generate the next cell states
	for {
		fmt.Printf("Generation(seed[%d]): %d\n", *gameSeed, generation)
		universe.Show()
		universe.Next()
		// don't generate anymore as all cells are dead
		if universe.Dead(){
			log.Println("universe is dead")
			os.Exit(1)
		}
		time.Sleep(500 * time.Millisecond) // wait for sometime
		generation++
	}
}

func main() {
	Evolve()
}
