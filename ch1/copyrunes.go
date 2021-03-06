package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	for {
		if r, _, err := input.ReadRune(); err == nil {
			if _, err := output.WriteRune(r); err != nil {
				panic("Error writing to STDOUT:" + err.Error())
			}
			if r == '\n' {
				output.Flush()
			}
		} else if err == io.EOF {
			break
		} else {
			panic("Error reading from STDIN:" + err.Error())
		}
	}
	output.Flush()
}
