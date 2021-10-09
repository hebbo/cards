package even

import (
	"fmt"
	"math/rand"
	"sort"
)

func Call() {
	// just making things more complicatec :D
	var sortedListOfNumbers []int = rand.Perm(10)
	sort.Ints(sortedListOfNumbers)

	for _, number := range sortedListOfNumbers {
		fmt.Println(number, "even: ", number%2 == 0)
	}
}
