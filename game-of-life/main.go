package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/lucasturci/everything-go/data-structures/matrix"
)

var n, m int
var mat matrix.Matrix[int]

func usage() {
	fmt.Printf("Usage: %s <n> [m], where \n\tn = number of rows\n\t"+
		"m = number of columns of game of life view.\nIf m is omitted, it will create a square view\n", os.Args[0])
	os.Exit(0)
}

func countNeighbours(i, j int) int {
	moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	ans := 0
	for _, move := range moves {
		p := i + move[0]
		q := j + move[1]
		if p < 0 || p >= n || q < 0 || q >= m {
			continue
		}
		if mat[p][q] == 1 {
			ans++
		}
	}
	return ans
}

func update() {
	// Rules
	// 1. If cell is alive and has less than 2 or more than 3 neighbours, kill it
	// 2. If cell is dead and has exactly 3 neighbours, revive it.
	mat2 := mat.Clone()

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt := countNeighbours(i, j)
			if mat[i][j] == 1 && (cnt < 2 || cnt > 3) {
				mat2[i][j] = 0
			} else if mat[i][j] == 0 && cnt == 3 {
				mat2[i][j] = 1
			}
		}
	}
	mat = mat2
}

var did bool = false

func print() {
	if did {
		for i := 0; i < n; i++ {
			fmt.Print("\033[1A\033[K")
		}
	}
	did = true

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			char := []byte{'.', 'O'}
			fmt.Print(string(char[mat[i][j]]))
		}
		fmt.Println()
	}
}

func main() {
	argc := len(os.Args)
	if argc <= 1 || argc > 3 {
		usage()
	}

	var err error
	n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	m = n
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

	// Initialize world
	mat = matrix.New[int](n, m)

	// initialize with seed
	for i := range mat {
		for j := 0; j < m; j++ {
			mat[i][j] = rand.Int() % 4
			if mat[i][j] < 3 {
				mat[i][j] = 0
			} else {
				mat[i][j] /= mat[i][j]
			}
		}
	}

	for {
		// print and update every half a second
		print()
		time.Sleep(250 * time.Millisecond)
		update()
	}

}
