package maps

import "fmt"

func Call() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#36bcde",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
