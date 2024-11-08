package main

import "fmt"

type englishBot struct {
}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// func (eb englishBot) getGreeting() string {
// 	//Very custom logic for generating an english greeting

// 	return "Hi There!";
// }

// we can ommit eb as we are not using it
func (englishBot) getGreeting() string {
	//Very custom logic for generating an english greeting

	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
