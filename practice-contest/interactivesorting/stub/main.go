package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//read query
	var fp *os.File
	var err error
	fp, err = os.Open("query.txt")
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	fmt.Println(scanner.Text())
}
