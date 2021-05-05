package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	// https://readme.so/it
	hasError := false
	var wg sync.WaitGroup
	reader := NewReader()

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

	log.Println("List of changed files")
	for _, status := range list {
		log.Println(" - ", status.File)
	}

	// decide which commands to execute
	commands := Commands(list, configurations)

	counter := 0

	// Execute commands in concurrency
	for _, cfg := range commands {
		wg.Add(1)
		go worker(cfg, &hasError, &counter, len(commands), &wg)
	}

	wg.Wait()

	if hasError {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
