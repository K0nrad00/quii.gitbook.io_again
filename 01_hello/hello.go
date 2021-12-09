package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name

}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("klkl", spanish))
	someRepeat("f")

}

func someRepeat(char string) {
	empty_str := ""
	for i := 0; i < 7; i++ {
		empty_str += char
	}
	fmt.Println(empty_str)
}
