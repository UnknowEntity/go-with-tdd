package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
)

func languagePrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return languagePrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
