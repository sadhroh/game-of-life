package main

import (
	"flag"
	"time"
)

// command line flags
var (
	gameSeed         = flag.Int64("seed", time.Now().UnixNano(), "starting seed")
	gameWorldRows    = flag.Int("rows", 40, "number of rows")
	gameWorldColumns = flag.Int("columns", 40, "number of columns")
)

func init() {
	flag.Parse()
}

func main() {
	universe := NewWorld(*gameSeed, *gameWorldRows, *gameWorldColumns)
	universe.Show()
	universe.Next()
}
