package main

import (
	"fmt"

	"github.com/radifan9/weeklytask-w8/internal/process"
)

func main() {
	fmt.Println("--- --- Weeklytask W8 --- ---")

	// --- --- No 1 --- ---
	runProcessNumber()
}

func runProcessNumber() {
	output := process.ProcessNumber(1, 2, 3, 4)
	if len(output) > 0 {
		fmt.Println(output)
	}

	// input nil test
	output2 := process.ProcessNumber()
	if len(output2) > 0 {
		fmt.Println(output2)
	}

	// daftar input kosong --> input terdapat zero value
	// zero valuenya int adalah 0
	output3 := process.ProcessNumber(0, 0)
	if len(output3) > 0 {
		fmt.Println(output3)
	}
}
