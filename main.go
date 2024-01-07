package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	//VERY CUSTOM LOGIC FOR GENERATING AN ENGLISH GREETING
	return "Hi There"
}

func (sb spanishBot) getGreeting() string {
	return "Hiola"
}