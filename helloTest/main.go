package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchHelloPrefix + name
	// }
func greetingPrefix(language string) (prefix string){

	switch language{
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

	fmt.Println(Hello("Jose", "French"))
}
