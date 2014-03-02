package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	nw := 0
	inword := false
	for {
		r, _, err := input.ReadRune()
		if err == nil {
			if unicode.IsSpace(r) {
				inword = false
			} else if ! inword {
				nw++
				inword = true
			}
		} else if err == io.EOF {
			break;
		} else {
			panic("Error reading from STDIN:" + err.Error())
		}
	}
	fmt.Printf("%d\n", nw)
}
