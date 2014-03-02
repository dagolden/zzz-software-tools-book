package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const NEWLINE = '\u000a'

func main() {
	input := bufio.NewReader(os.Stdin)
	nl := 0
	for {
		r, _, err := input.ReadRune()
		if err == nil {
			if r == NEWLINE {
				nl++
			}
		} else if err == io.EOF {
			break;
		} else {
			panic("Error reading from STDIN:" + err.Error())
		}
	}
	fmt.Printf("%d\n", nl)
}
