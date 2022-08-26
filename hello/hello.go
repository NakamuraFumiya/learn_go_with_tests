package main

import "fmt"

const englistHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}
	return englistHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}