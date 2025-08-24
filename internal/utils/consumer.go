package utils

import "log"

// <-chan : only receive
func FetchConsumer(urlResults <-chan string) {
	for result := range urlResults {
		log.Println(result)
	}

}
