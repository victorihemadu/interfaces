package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting (b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY CUSTOM LOGIC FOR GENERATING AN ENGLISH GREETING
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hiola"
}

type user struct {
	name string
}
type Bot interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float32
	respondToUser(user) string
}