package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	se := os.Stderr

	fmt.Fprintln(se, "stub start")

	fmt.Fprintln(se, "read ./query.fifo")
	fmt.Fprintln(se, "write ./result.fifo")
	fp, err := os.OpenFile("./query.fifo", os.O_RDONLY|os.O_CREATE, os.ModeNamedPipe)
	o, oe := os.OpenFile("./result.fifo", os.O_RDWR|os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if oe != nil {
		panic(oe)
	}
	defer o.Close()

	t := makeTable()

	b := make([]byte, 1024)
	for {
		n, err := fp.Read(b)
		if err != io.EOF && err != nil {
			panic(err)
		}
		if n != 0 {
			fmt.Fprintf(se, "query: %s", string(b))
			lw := t[b[2]]
			rw := t[b[4]]
			var a string
			if lw < rw {
				a = "<"
			} else {
				a = ">"
			}
			_, e := fmt.Fprintln(o, a)
			fmt.Fprintln(se, "wrote")
			if e != nil {
				panic(e)
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func makeTable() map[byte]int {
	return map[byte]int{
		[]byte("A")[0]: 1,
		[]byte("B")[0]: 0,
		[]byte("C")[0]: 23,
		[]byte("D")[0]: 9,
		[]byte("E")[0]: 4,
		[]byte("F")[0]: 5,
		[]byte("G")[0]: 6,
		[]byte("H")[0]: 3,
		[]byte("I")[0]: 8,
		[]byte("J")[0]: 14,
		[]byte("K")[0]: 12,
		[]byte("L")[0]: 7,
		[]byte("M")[0]: 15,
		[]byte("N")[0]: 13,
		[]byte("O")[0]: 16,
		[]byte("P")[0]: 2,
		[]byte("Q")[0]: 20,
		[]byte("R")[0]: 11,
		[]byte("S")[0]: 19,
		[]byte("T")[0]: 22,
		[]byte("U")[0]: 10,
		[]byte("V")[0]: 18,
		[]byte("W")[0]: 17,
		[]byte("Z")[0]: 21,
	}
}
