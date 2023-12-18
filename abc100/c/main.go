package main

import (
	"bufio"
	"fmt"
	"os"
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

	outPut:= inputLine
	// 標準出力で解答を出す
	fmt.Println(outPut)
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}

