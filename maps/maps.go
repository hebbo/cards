package maps

import "fmt"

func Call() {
	empty := make(map[int]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#36bcde",
	}

	empty[10] = "#ffffff"
	delete(empty, 10)

	fmt.Println(colors)
	fmt.Println(empty)
}
