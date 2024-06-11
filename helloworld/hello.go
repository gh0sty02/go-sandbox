package main

const (
	spanish            = "spanish"
	french             = "french"
	englishHelloPrefix = "hello "
	spanishHelloPrefix = "hola "
	frenchHelloPrefix  = "bonjour "
)

func hello(name, language string) string {
	if name == "" {
		return "hello world"
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

	return
}

func main() {
	println(hello("Chris", "spanish"))
}
