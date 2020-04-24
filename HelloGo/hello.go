package main

import "fmt"

const HelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const Spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == Spanish {
		return SpanishHelloPrefix + name
	}
	return HelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
