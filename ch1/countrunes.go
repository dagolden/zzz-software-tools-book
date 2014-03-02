package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	nc := 0
	for {
		if _, _, err := input.ReadRune(); err != nil {
			if err == io.EOF {
				break
			}
			panic("Error reading from STDIN:" + err.Error())
		}
		nc++
	}
	fmt.Printf("%d\n",nc)
}
