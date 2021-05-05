package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func worker(cfg Configuration, hasError *bool, counter *int, numberOfJobs int, wg *sync.WaitGroup) {
	defer wg.Done()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command(cfg.Command.Name, cfg.Command.Arg)
	cmd.Dir = fmt.Sprintf("%s/%s", path, cfg.Directory)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	*counter++
	counterString := fmt.Sprintf("[%d/%d]", *counter, numberOfJobs)

	if err != nil {
		*hasError = true
		log.Printf(ErrorColor, fmt.Sprintf("✗ %s %s", counterString, cfg.Name))
		//fmt.Println(out.String())
		//fmt.Println(stderr.String())
	} else {
		log.Printf(NoticeColor, fmt.Sprintf("✓ %s %s", counterString, cfg.Name))
	}
}
