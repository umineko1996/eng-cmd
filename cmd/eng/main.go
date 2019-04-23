package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var out io.Writer = os.Stdout
var EnoguAA = loadAA("data/enogu_logo_AA.txt")

func main() {
	fmt.Fprintln(out, "\x1b[2J")
	for i := 50; i > 0; i-- {
		for j, line := range strings.Split(EnoguAA, "\r\n") {
			fmt.Fprintf(out, "\x1b[%d;%dH\x1bK", j+5, i)
			fmt.Println(line, " ")
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func loadAA(filename string) string {
	b := MustAsset(filename)
	return string(b)
}
