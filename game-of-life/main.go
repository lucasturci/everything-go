package main

/*

TODO:

- make game work with different rules
- make 1st positional argument another option to pass pattern
- change window title to be the name of the pattern
- give option to user to press esc to quit program
- better error handling
*/

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lucasturci/everything-go/data-structures/matrix"
	"github.com/lucasturci/everything-go/game-of-life/patternfetcher"
)

var n, m int
var patternArg string

const CELL_SIZE = 20
const TICK_SPEED = 10
const PATTERN_COUNT = 20 // number of random patterns to select to the user

type Game struct {
	world     matrix.Matrix[int]
	cellSize  int
	rows      int
	cols      int
	lastTick  int
	tickSpeed int // Update every N frames
}

func createRandomPattern(rows, cols int) matrix.Matrix[int] {
	world := matrix.New[int](rows, cols)
	// Initialize with random cells
	for i := range world {
		for j := 0; j < cols; j++ {
			world[i][j] = rand.Int() % 4
			if world[i][j] < 3 {
				world[i][j] = 0
			} else {
				world[i][j] = 1
			}
		}
	}
	return world
}

func NewGame(rows, cols int, pattern matrix.Matrix[int]) *Game {
	g := &Game{
		world:     matrix.New[int](rows, cols),
		cellSize:  CELL_SIZE,
		rows:      rows,
		cols:      cols,
		tickSpeed: TICK_SPEED,
	}

	// fill the world with pattern
	// Calculate starting position to center the pattern
	startRow := (rows - pattern.SizeRows()) / 2
	startCol := (cols - pattern.SizeCols()) / 2

	// Copy pattern into center of world
	for i := 0; i < pattern.SizeRows(); i++ {
		for j := 0; j < pattern.SizeCols(); j++ {
			g.world[startRow+i][startCol+j] = pattern[i][j]
		}
	}
	return g
}

func (g *Game) Update() error {
	// Update simulation every N frames
	g.lastTick++
	if g.lastTick >= g.tickSpeed {
		g.updateWorld()
		g.lastTick = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 40, 255}) // Dark gray background

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.world[i][j] == 1 {
				vector.DrawFilledRect(
					screen,
					float32(j*g.cellSize),
					float32(i*g.cellSize),
					float32(g.cellSize-1),
					float32(g.cellSize-1),
					color.RGBA{240, 240, 240, 255},
					false,
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cols * g.cellSize, g.rows * g.cellSize
}

func (g *Game) countNeighbours(i, j int) int {
	moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	ans := 0
	for _, move := range moves {
		p := i + move[0]
		q := j + move[1]
		if p < 0 || p >= g.rows || q < 0 || q >= g.cols {
			continue
		}
		if g.world[p][q] == 1 {
			ans++
		}
	}
	return ans
}

func (g *Game) updateWorld() {
	newWorld := g.world.Clone()

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			cnt := g.countNeighbours(i, j)
			if g.world[i][j] == 1 && (cnt < 2 || cnt > 3) {
				newWorld[i][j] = 0
			} else if g.world[i][j] == 0 && cnt == 3 {
				newWorld[i][j] = 1
			}
		}
	}
	g.world = newWorld
}

func main() {
	flag.IntVar(&n, "n", 40, "number of rows in world matrix")
	flag.IntVar(&m, "m", 40, "number of columns in world matrix")
	flag.StringVar(&patternArg, "pattern", "", "seed pattern. From this index of patterns: https://conwaylife.com/patterns/")

	flag.Parse()

	var pattern matrix.Matrix[int]
	if len(patternArg) > 0 { // user passed pattern, only fetch that
		if patternArg == "random" {
			pattern = createRandomPattern(n, m)
		} else {
			pattern = patternfetcher.ParsePattern(patternArg)
		}
	} else {
		fmt.Print("Selecting random patterns for you to choose.\nFull list of patterns you can choose from is available at https://conwaylife.com/patterns/\n")
		patterns := patternfetcher.ScrapePatterns()
		patterns["random"] = createRandomPattern(n, m)

		fmt.Printf("Loaded %v patterns, here are the options:\n", len(patterns))
		i := 0
		for name := range patterns {
			fmt.Printf("%-40s", name)
			if i > 0 && i%3 == 0 {
				fmt.Println()
			}
			i++
		}

		var choice string
		for {
			fmt.Printf("Choose one: \n")
			fmt.Scanf("%s", &choice)
			_, ok := patterns[choice]
			if ok {
				fmt.Println("Nice choice!")
				break
			}
			fmt.Println("That's not an option, try again")
		}

		pattern = patterns[choice]
	}

	if n < 4 || m < 4 {
		fmt.Println("Please input dimensions greater than 3 for better visualization")
		return
	}

	n = max(n, pattern.SizeRows())
	m = max(m, pattern.SizeCols())
	ebiten.SetWindowSize(m*CELL_SIZE, n*CELL_SIZE)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(NewGame(n, m, pattern)); err != nil {
		log.Fatal(err)
	}
}
