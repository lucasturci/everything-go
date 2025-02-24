package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lucasturci/everything-go/data-structures/matrix"
)

var n, m int
var mat matrix.Matrix[int]

const CELL_SIZE = 20
const TICK_SPEED = 10

type Game struct {
	world     matrix.Matrix[int]
	cellSize  int
	rows      int
	cols      int
	lastTick  int
	tickSpeed int // Update every N frames
}

func usage() {
	fmt.Printf("Usage: %s <n> [m], where \n\tn = number of rows\n\t"+
		"m = number of columns of game of life view.\nIf m is omitted, it will create a square view\n", os.Args[0])
	os.Exit(0)
}

func NewGame(rows, cols int) *Game {
	g := &Game{
		world:     matrix.New[int](rows, cols),
		cellSize:  CELL_SIZE, // Adjust this to change cell size
		rows:      rows,
		cols:      cols,
		tickSpeed: 10, // Adjust this to change simulation speed
	}

	// Initialize with random cells
	for i := range g.world {
		for j := 0; j < cols; j++ {
			g.world[i][j] = rand.Int() % 4
			if g.world[i][j] < 3 {
				g.world[i][j] = 0
			} else {
				g.world[i][j] = 1
			}
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
	argc := len(os.Args)
	if argc <= 1 || argc > 3 {
		usage()
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	m := n
	if argc == 3 {
		m, err = strconv.Atoi(os.Args[2])
		if err != nil {
			usage()
		}
	}

	if n < 4 || m < 4 {
		fmt.Println("Please input dimensions greater than 3 for better visualization")
		return
	}

	ebiten.SetWindowSize(m*CELL_SIZE, n*CELL_SIZE)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(NewGame(n, m)); err != nil {
		log.Fatal(err)
	}
}
