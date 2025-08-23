package process

import (
	"errors"
	"fmt"
)

func ProcessNumber(input ...int) []int {
	var err error

	// General Error
	defer func() {
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// Panic Error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// Check if its a nil or not (General Error)
	if input == nil {
		err = errors.New("no data provided")
	}

	// Make an empty slice
	newSlice := make([]int, 0)
	for _, value := range input {
		// Check if the list is empty (0)
		if value == 0 {
			panic("empty list provided")
		}

		newSlice = append(newSlice, value*2)
	}

	return newSlice
}
