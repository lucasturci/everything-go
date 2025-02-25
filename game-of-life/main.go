package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/lucasturci/everything-go/algorithms/encoding"
	"github.com/lucasturci/everything-go/data-structures/matrix"
	"github.com/lucasturci/everything-go/data-structures/tuple"
)

var n, m int
var pat string

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

func parsePattern(patternURI string) matrix.Matrix[int] {
	log.Printf("Parsing pattern %s\n", patternURI)
	// Fetch the patterns page
	resp, err := http.Get("https://conwaylife.com/patterns/" + patternURI + ".rle")
	if err != nil {
		log.Printf("Failed to fetch pattern %s: ", patternURI, err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response line by line
	scanner := bufio.NewScanner(resp.Body)
	code := ""
	re := regexp.MustCompile(`^x=(\d+),y=(\d+),rule=B3\/S23`)
	var w, h int
	// decode using https://conwaylife.com/wiki/Run_Length_Encoded as reference
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "") // remove all spaces
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 {
			x, err := strconv.Atoi(matches[0][1])
			if err != nil {
				panic("regex didn't match an integer")
			}
			w = x
			y, err := strconv.Atoi(matches[0][2])
			h = y
			if err != nil {
				panic("regex didn't match an integer")
			}
		} else {
			code += line
		}
	}
	if w == 0 || h == 0 {
		return nil
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading pattern file: %v\n", err)
		return nil
	}

	pattern := matrix.New[int](h+1, w+1)
	codeEnd := strings.Index(code, "!")
	if codeEnd != -1 {
		code = code[:codeEnd]
	} else {
		log.Printf("Did not find ! at the end of " + patternURI + " RLE code\n")
	}

	rle := encoding.NewRLE(true)
	decodedBytes, err := rle.Decode([]byte(code))
	if err != nil {
		log.Printf("error: %v wrong happened when decoding %s - code %s\n", patternURI, err, code)
		return nil
	}
	log.Printf("Building matrix of pattern %s with size (%v, %v)\n", patternURI, h, w)
	decoded := string(decodedBytes)
	for i, row := range strings.Split(decoded, "$") {
		for j, c := range string(row) {
			if c == 'o' {
				pattern[i][j] = 1
			}
		}
	}

	return pattern
}

func scrapePatterns() map[string]matrix.Matrix[int] {
	patterns := make(map[string]matrix.Matrix[int])

	// Fetch the patterns page
	resp, err := http.Get("https://conwaylife.com/patterns/")
	if err != nil {
		log.Printf("Failed to fetch patterns: %v", err)
		return patterns
	}
	defer resp.Body.Close()

	// Read the entire body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
		return patterns
	}
	re := regexp.MustCompile(`href="\/patterns\/(.*)\.rle"`)
	matches := re.FindAllSubmatch(body, -1)
	fmt.Printf("Found %v patterns\n", len(matches))

	rand.Shuffle(len(matches), func(i, j int) {
		matches[i], matches[j] = matches[j], matches[i]
	})

	matches = matches[:10]

	patternChan := make(chan tuple.Pair[string, matrix.Matrix[int]])
	for _, match := range matches {
		url := match[1]
		go func(url string) {
			res := parsePattern(string(url))
			patternChan <- tuple.Pair[string, matrix.Matrix[int]]{
				First:  url,
				Second: res,
			}
		}(string(url))
	}

	for range matches {
		keyValuePair := <-patternChan
		if keyValuePair.Second == nil {
			continue
		}
		patterns[keyValuePair.First] = keyValuePair.Second
	}
	close(patternChan)

	return patterns
}

func main() {
	flag.IntVar(&n, "n", 40, "number of rows in world matrix")
	flag.IntVar(&m, "m", 40, "number of columns in world matrix")
	flag.StringVar(&pat, "pattern", "", "seed pattern. From this index of patterns: https://conwaylife.com/patterns/")

	flag.Parse()

	var pattern matrix.Matrix[int]
	if len(pat) > 0 { // user passed pattern, only fetch that
		if pat == "random" {
			pattern = createRandomPattern(n, m)
		} else {
			pattern = parsePattern(pat)
		}
	} else {
		fmt.Print("Selecting random patterns for you to choose.\nFull list of patterns you can choose from is available at https://conwaylife.com/patterns/\n")
		patterns := scrapePatterns()
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
