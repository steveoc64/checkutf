package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("USAGE: checkfile filename")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		l := scanner.Text()
		u := strings.ToValidUTF8(l, "")
		if u != l {
			println("line", i, "has some invalid UTF8 chars")
			println(l)
		}
	}
	println("got lines", i)
}
