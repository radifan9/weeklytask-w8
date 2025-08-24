package main

import (
	"fmt"
	"sync"

	"github.com/radifan9/weeklytask-w8/internal/process"
	"github.com/radifan9/weeklytask-w8/internal/utils"
)

func main() {
	fmt.Println("--- --- Weeklytask W8 --- ---")

	// --- --- No 1 --- ---
	// runProcessNumber()

	// --- --- No 2 --- ---
	// runWebFetcher()

	// --- --- No 3 --- ---
	runUserManager()
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

func runWebFetcher() {
	// List of URLs to fetch (simulated)
	urls := []string{
		"www.google.com",
		"www.facebook.com",
		"www.spotify.com",
		"www.koda.com",
		"www.microsoft.com",
	}

	resultsChan := make(chan string)

	var fetchingWg sync.WaitGroup

	// Fetching routine
	for _, url := range urls {
		fetchingWg.Add(1)
		go utils.WebFetcher(url, &fetchingWg, resultsChan)
	}

	// Consumer routine
	var consumerWg sync.WaitGroup
	consumerWg.Add(1)
	go func() {
		utils.FetchConsumer(resultsChan)
		// When FetchConsumer is finishes (channel is closed and all results are consumed) wg.Done is called

		consumerWg.Done()
	}()

	fetchingWg.Wait()
	close(resultsChan)
	consumerWg.Wait()
}

func handleUserAddition(usersDB *utils.UserManager, id, name string) {
	success, err := usersDB.AddUser(id, name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(success)
	}
}

func handleUserGetter(usersDB *utils.UserManager, id string) {
	success, err := usersDB.GetUser(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(success)
	}
}

func runUserManager() {
	// Initialize UserManager instance
	usersDB := utils.NewUserManager()

	// User creation test cases
	fmt.Println("--- --- Setting the User Test --- ---")
	handleUserAddition(&usersDB, "01", "Radif")
	handleUserAddition(&usersDB, "02", "Opet")
	handleUserAddition(&usersDB, "03", "Yusuf")
	handleUserAddition(&usersDB, "01", "Radif")

	// User retrieval test cases
	fmt.Println("--- --- Getting the User Test --- ---")
	handleUserGetter(&usersDB, "01")
	handleUserGetter(&usersDB, "04")
}
