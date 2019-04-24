package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var out io.Writer = os.Stdout
var enoguAA = loadAA("data/enogu_logo_AA.txt")

func main() {
	printLogoAA()
}

func printLogoAA() {
	fmt.Fprintln(out, "\x1b[2J")
	for i := 50; i > 0; i-- {
		for j, line := range strings.Split(enoguAA, "\r\n") {
			fmt.Fprintf(out, "\x1b[%d;%dH\x1bK", j+5, i)
			if j < 18 {
				fmt.Println(line, " ")
				continue
			}
			x := 0
			for _, r := range line {
				var color int
				if isWidth(r) {
					x += 2
				} else {
					x++
				}
				switch {
				case x < 11:
					color = 0
				case x < 21:
					color = 31
				case x < 31:
					color = 35
				case x < 43:
					color = 36
				case x < 54:
					color = 32
				default:
					color = 33
				}
				fmt.Fprintf(out, "\x1b[%dm%s", color, string(r))
			}
			fmt.Println(" ")
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func loadAA(filename string) string {
	b := MustAsset(filename)
	return string(b)
}

func isWidth(r rune) bool {
	if r <= 0xF0 {
		// ASCII: size 1
		return false
	}
	return true
}
