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
}

func newSorter(w io.Writer, r io.Reader, lw io.Writer) sorter {
	return sorter{
		bufio.NewWriter(w),
		bufio.NewScanner(r),
		&lw,
	}
}

func (s *sorter) execute() {
	n, q := s.readNQ()
	fmt.Fprintf(*s.logWriter, "read N: %d, Q: %d\n", n, q)

	array := s.makeUnsorted(n)
	fmt.Fprintf(*s.logWriter, "initial array: %q\n", array)

	//sort first two pairs
	r1 := s.query(array[0], array[1])
	if r1 == ">" {
		t := array[0]
		array[0] = array[1]
		array[1] = t
	}
	r2 := s.query(array[2], array[3])
	if r2 == ">" {
		t := array[2]
		array[2] = array[3]
		array[3] = t
	}
	fmt.Println(array)
}

func (s *sorter) readNQ() (int, int) {
	s.scanner.Scan()
	t := []rune(s.scanner.Text())
	n, _ := strconv.Atoi(string(t[0]))
	q, _ := strconv.Atoi(string(t[2]))
	return n, q
}

func (s *sorter) query(l, r string) string {
	q := "? " + l + " " + r + "\n"
	fmt.Fprintln(*s.logWriter, q)
	s.writer.WriteString(q)
	s.writer.Flush()

	s.scanner.Scan()
	a := s.scanner.Text()
	fmt.Fprintln(*s.logWriter, a)
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
