package patternfetcher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucasturci/everything-go/algorithms/encoding"
	"github.com/lucasturci/everything-go/data-structures/matrix"
	"github.com/lucasturci/everything-go/data-structures/tuple"
)

func ParsePattern(patternURI string) matrix.Matrix[int] {
	log.Printf("Parsing pattern %s\n", patternURI)
	// Fetch the patterns page
	resp, err := http.Get("https://conwaylife.com/patterns/" + patternURI + ".rle")
	if err != nil {
		log.Printf("Failed to fetch pattern %s: %w", patternURI, err)
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

func ScrapePatterns() map[string]matrix.Matrix[int] {
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
			res := ParsePattern(string(url))
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
