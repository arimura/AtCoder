package main

import "fmt"

func main() {
	var a, b, c int
	var d string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	fmt.Printf("%d %s\n", a+b+c, d)
}
