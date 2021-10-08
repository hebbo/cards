package main

import "fmt"

func main() {
	d := newDeck()
	fmt.Println("writing to file...")
	fmt.Println(d.saveToFile("udemy_golang"))
}
