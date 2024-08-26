package main

import (
	"fmt"
)

const (
	spanish    = "Spanish"
	french     = "French"
	lithuanian = "Lithuanian"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	lithuanianHelloPrefix = "Labas, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case lithuanian:
		prefix = lithuanianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Ernestas", ""))
}
