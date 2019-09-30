package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type sorter struct {
	writer    *bufio.Writer
	scanner   *bufio.Scanner
	logWriter *io.Writer
	cnt       int
}

func newSorter(w io.Writer, r io.Reader, lw io.Writer) sorter {
	return sorter{
		bufio.NewWriter(w),
		bufio.NewScanner(r),
		&lw,
		0,
	}
}

func (s *sorter) execute() {
	n, q := s.readNQ()
	fmt.Fprintf(*s.logWriter, "read N: %d, Q: %d\n", n, q)

	array := s.makeUnsorted(n)
	fmt.Fprintf(*s.logWriter, "initial array: %q\n", array)

	//sort first two pairs
	if !s.isGreater(array[0], array[1]) {
		array[0], array[1] = array[1], array[0]
	}

	if !s.isGreater(array[2], array[3]) {
		array[2], array[3] = array[3], array[2]
	}

	//compare greater elements of first two pairs
	if !s.isGreater(array[1], array[3]) {
		array[0], array[1], array[2], array[3] = array[2], array[3], array[0], array[1]
	}

	fmt.Fprintf(*s.logWriter, "result of 3rd querying: %q", array)

	//array = [a,b,c,d,e] with a < b < d and c < d
	//insert e int [a,b,d]
	var tmpArray []string
	var cElement = array[2]
	if s.isGreater(array[4], array[1]) { //e < b
		//e < b
		if s.isGreater(array[4], array[0]) { //e < a
			//e < a
			//e, a, b, d (c: not determined)
			tmpArray = []string{array[4], array[0], array[1], array[3]}
		} else {
			//a < e
			//a, e, b ,d (c: not determined)
			tmpArray = []string{array[0], array[4], array[1], array[3]}
		}
	} else {
		//b < e
		if s.isGreater(array[4], array[3]) { //e < d
			//e < d
			//a, b, e ,d (c: not determined)
			tmpArray = []string{array[0], array[1], array[4], array[3]}
		} else {
			//d< e
			//a, b, d, e (c: not determined)
			tmpArray = []string{array[0], array[1], array[3], array[4]}
		}
	}
	var sortedArray []string
	if s.isGreater(cElement, tmpArray[1]) {
		if s.isGreater(cElement, tmpArray[0]) {
			sortedArray = []string{cElement, tmpArray[0], tmpArray[1], tmpArray[2], tmpArray[3]}
		} else {
			sortedArray = []string{tmpArray[0], cElement, tmpArray[1], tmpArray[2], tmpArray[3]}
		}
	} else {
		if s.isGreater(cElement, tmpArray[2]) {
			sortedArray = []string{tmpArray[0], tmpArray[1], cElement, tmpArray[2], tmpArray[3]}
		} else {
			sortedArray = []string{tmpArray[0], tmpArray[1], tmpArray[2], cElement, tmpArray[3]}
		}
	}

	fmt.Fprintf(*s.logWriter, "total query count: %d\n", s.cnt)
	fmt.Fprintf(*s.logWriter, "sorted: %q\n", sortedArray)
}

func (s *sorter) readNQ() (int, int) {
	s.scanner.Scan()
	t := []rune(s.scanner.Text())
	n, _ := strconv.Atoi(string(t[0]))
	q, _ := strconv.Atoi(string(t[2]))
	return n, q
}

func (s *sorter) isGreater(l, r string) bool {
	return s.query(l, r) == "<"
}

func (s *sorter) query(l, r string) string {
	q := "? " + l + " " + r + "\n"
	fmt.Fprintln(*s.logWriter, q)
	s.writer.WriteString(q)
	s.writer.Flush()

	s.scanner.Scan()
	a := s.scanner.Text()
	fmt.Fprintln(*s.logWriter, a)
	s.cnt++
	return a
}

func (s *sorter) makeUnsorted(n int) []string {
	str := make([]string, n)
	for index := 0; index < n; index++ {
		c := index + int('A')
		str[index] = string(c)
	}
	return str
}

func main() {
	logWriter := os.Stderr
	defer logWriter.Close()

	//writer
	writer, err := os.OpenFile("./sorter_to_stub.fifo", os.O_WRONLY|os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	//reader
	reader, err := os.OpenFile("./stub_to_sorter.fifo", os.O_RDONLY|os.O_CREATE, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	sorter := newSorter(writer, reader, logWriter)
	sorter.execute()
}

func makeUnsorted(n int) []string {
	s := make([]string, n)
	for index := 0; index < n; index++ {
		c := index + int('A')
		s[index] = string(c)
	}
	return s
}

func query(w io.Writer, reader io.Reader, l, r string) string {
	bf := bufio.NewWriter(w)
	bf.WriteString("? " + l + " " + r + "\n")
	bf.Flush()

	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return scanner.Text()
}
