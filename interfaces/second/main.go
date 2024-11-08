package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

type hindiBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(hb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

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

func (hindiBot) getGreeting() string {
	return "Namaste!"
}
