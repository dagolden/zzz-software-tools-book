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

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	stops := setStops()
	col, newcol := 1, 1
	var r rune
	var err error
	for {
		newcol = col
		r, err = getC(input)
		for r == SPACE && err == nil {
			newcol++
			if isStop(stops, newcol) {
				putC(output, TAB)
				col = newcol
			}
			r, err = getC(input)
		}
		for col < newcol {
			putC(output, SPACE)
			col++
		}
		if err != io.EOF {
			putC(output, r)
			if r == NEWLINE {
				output.Flush()
				col = 1
			} else {
				col++
			}
		} else {
			break
		}
	}
	output.Flush()
}
