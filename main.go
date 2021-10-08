package main

func main() {
	mydeck := newDeckFromFile("udemy_golang")

	// fmt.Println("before shuffling...")
	// mydeck.print()

	mydeck.shuffle()

	// fmt.Println("after shuffling...")
	mydeck.print()
}
