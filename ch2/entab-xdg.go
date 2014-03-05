package main

import (
	"bufio"
	"io"
	"os"
)

const (
	TABSPACE = 4
	MAXLINE  = 1000
	TAB      = '\t'
	SPACE    = ' '
	NEWLINE  = '\n'
)

func setStops() []bool {
	tabStops := make([]bool, MAXLINE+1)
	for i := range tabStops {
		tabStops[i] = (i%TABSPACE == 1)
	}
	return tabStops
}

func isStop(stops []bool, n int) bool {
	if n > MAXLINE {
		return true
	} else {
		return stops[n]
	}
}

func getC(input *bufio.Reader) (rune, error) {
	r, _, err := input.ReadRune()
	if err == nil || err == io.EOF {
		return r, err
	} else {
		panic("Error reading from STDIN:" + err.Error())
	}
}

func putC(output *bufio.Writer, r rune) {
	if _, err := output.WriteRune(r); err != nil {
		panic("Error writing to STDOUT:" + err.Error())
	}
}

func putSpaces(output *bufio.Writer, stops []bool, col, ahead int) int {
	// output tabs
	// NB if this runs only when col + 1 < ahead, then it won't output a TAB
	// when a single space would be acceptable; but this deviates from
	// the Software Tools book
	for i := col; i < ahead; i++ {
		if isStop(stops, i+1) {
			putC(output, TAB)
			col = i + 1
		}
	}

	// output any remaining spaces
	for i := col; i < ahead; i++ {
		putC(output, SPACE)
	}

	return ahead
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	stops := setStops()
	col, ahead := 1, 1
	for {
		r, err := getC(input)
		// ought to check for col or ahead > MAXLINE
		if err == io.EOF {
			break
		} else if r == SPACE {
			ahead++
		} else {
			if col < ahead {
				col = putSpaces(output, stops, col, ahead)
			}
			putC(output, r)
			if r == NEWLINE {
				col = 1
				output.Flush()
			} else {
				col++
			}
			ahead = col
		}
	}
	putSpaces(output, stops, col, ahead)
	output.Flush()
}
