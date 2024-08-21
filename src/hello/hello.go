package main

import "fmt"

func main() {
	fmt.Println(printHello("Brian"))
}

const englishHelloPrefix = "Hello"

func printHello(name string) string {

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + ", " + name
}
