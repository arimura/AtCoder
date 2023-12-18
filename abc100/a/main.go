package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	const max = 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
}

func main() {
	// 問題のcaseは標準入力で受け取る
	inputLine := NextLine(sc)
	inputs := strings.Split(inputLine, " ")

	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])

	if a > 8 || b > 8 {
		fmt.Println(":(")
	} else {
		fmt.Println("Yay!")
	}
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}
