package main

import (
	"io"
	"os"
)

func main() {
	b := make([]byte, 1)
	for {
		if _, err := os.Stdin.Read(b); err == nil {
			if _, err := os.Stdout.Write(b); err != nil {
				panic("Error writing to STDOUT:" + err.Error())
			}
		} else if err == io.EOF {
			return
		} else {
			panic("Error reading from STDIN:" + err.Error())
		}
	}
}
