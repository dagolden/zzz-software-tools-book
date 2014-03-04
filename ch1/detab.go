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
	NEWLINE  = '\n'
)

func setStops() []bool {
	tabStops := make([]bool, 1000)
	for i := range tabStops {
		tabStops[i] = (i%TABSPACE == 1)
	}
	return tabStops
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
	col := 1
	for {
		r, err := getC(input)
		if err == io.EOF {
			output.Flush()
			break
		} else {
			if r == TAB {
				// fill in with spaces
				for {
					putC(output, ' ')
					col++
					if col > MAXLINE || stops[col] {
						break
					}
				}
			} else if r == NEWLINE {
				col = 1
				putC(output, r)
				output.Flush()
			} else {
				putC(output, r)
			}
		}
	}
}
