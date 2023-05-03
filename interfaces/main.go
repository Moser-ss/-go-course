package main

type bot interface {
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

func (eb englishBot) getGreeting() string {
	return "Hello there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola mucacho"
}

func printGreeting(b bot) {
	println(b.getGreeting())
}
