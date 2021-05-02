package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	hasError := false
	var wg sync.WaitGroup
	reader := NewReader()

	log.Println("Inizio")

	// READ configurations
	configurations, err := reader.Configurations()
	if err != nil {
		panic(err)
	}

	// READ git status
	list, err := reader.Status()
	if err != nil {
		panic(err)
	}

	// decide which commands to execute
	commands := Commands(list, configurations)

	// Execute commands in concurrency
	for _, cfg := range commands {
		wg.Add(1)
		go worker(cfg, &hasError, &wg)
	}

	wg.Wait()
	log.Println("Fine")

	if hasError {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
