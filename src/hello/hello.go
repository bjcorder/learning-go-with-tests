package main

import "fmt"

func main() {
	fmt.Println(printHello("Brian"))
}

func printHello(name string) string {
	return "Hello, " + name
}
