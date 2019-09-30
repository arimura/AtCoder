package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	se := os.Stderr

	n := flag.String("N", "5", "number of array")
	q := flag.String("Q", "7", "number of query")

	fmt.Fprintln(se, "stub start")

	fmt.Fprintln(se, "read ./sorter_to_stub.fifo")
	fmt.Fprintln(se, "write ./stub_to_sorter.fifo")
	fp, err := os.OpenFile("./sorter_to_stub.fifo", os.O_RDONLY|os.O_CREATE, os.ModeNamedPipe)
	o, oe := os.OpenFile("./stub_to_sorter.fifo", os.O_WRONLY|os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if oe != nil {
		panic(oe)
	}
	defer o.Close()

	fmt.Fprintln(o, *n+" "+*q)
	fmt.Fprintln(se, *n+" "+*q)

	t := makeTable()

	b := make([]byte, 1024)
	for {
		n, err := fp.Read(b)
		if err != io.EOF && err != nil {
			panic(err)
		}
		if n != 0 {
			in := string(b)
			fmt.Fprintf(se, "in: %s", in)
			if ([]rune(in))[0] == '!' {
				fmt.Fprintf(se, "got answer!!!\n")
				os.Exit(0)
			}

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
