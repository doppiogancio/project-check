package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func worker(cfg Configuration, hasError *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command(cfg.Command.Name, cfg.Command.Arg)
	cmd.Dir = cfg.Directory

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		*hasError = true
		log.Println("✗", cfg.Name, "ended with errors")
		fmt.Println(out.String())
		fmt.Println(stderr.String())
	} else {
		log.Println("✓", cfg.Name, "ended correctly")
	}
}
