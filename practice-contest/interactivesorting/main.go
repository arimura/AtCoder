package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := []rune(scanner.Text())
	n, _ := strconv.Atoi(string(t[0]))
	fmt.Printf("input N: %d\n", n)
	array := makeUnsorted(n)
	fmt.Printf("initial string: %s\n", array)

	//use fifo
	q, err := os.OpenFile("./query.fifo", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer q.Close()

	result, err := os.OpenFile("./result.fifo", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	//sort first two pairs
	r1 := query(q, result, array[0], array[1])
	fmt.Println(r1)
	// if array[0] > array[1] {
	// 	t := array[0]
	// 	array[0] = array[1]
	// 	array[1] = t
	// }
	// if array[2] > array[3] {
	// 	t := array[2]
	// 	array[2] = array[3]
	// 	array[3] = t
	// }

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
	w.Write([]byte("? " + l + " " + r))

	b := make([]byte, 1024)
	reader.Read(b)
	return string(b)
}
