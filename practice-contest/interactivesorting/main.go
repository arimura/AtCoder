package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := []rune(scanner.Text())
	n, _ := strconv.Atoi(string(t[0]))
	fmt.Printf("input N: %d\n", n)
	fmt.Printf("initial string: %s\n", makeUnsorted(n))

}

func makeUnsorted(n int) string {
	var s string
	for index := int('A'); index < int('A')+n; index++ {
		s += string(index)
	}
	return s
}
