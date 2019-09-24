package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("stub start")

	var fp *os.File
	var err error

	fmt.Println("read ./query.fifo")
	fp, err = os.OpenFile("./query.fifo", os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	t := makeTable()

	b := make([]byte, 1024)
	for {
		n, err := fp.Read(b)
		if err != io.EOF && err != nil {
			panic(err)
		}
		if n != 0 {
			fmt.Printf("query: %s", string(b))
			lw := t[b[2]]
			rw := t[b[4]]
			var a string
			if lw < rw {
				a = "<"
			} else {
				a = ">"
			}
			fmt.Println(a)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func makeTable() map[byte]int {
	return map[byte]int{
		[]byte("A")[0]: 1,
		[]byte("B")[0]: 0,
		[]byte("C")[0]: 2,
	}
}
