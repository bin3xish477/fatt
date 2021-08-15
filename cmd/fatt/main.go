package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/alexflint/go-arg"
	"github.com/binexisHATT/fatt/models"
	"github.com/binexisHATT/fatt/strings"
)

const (
	red     = "\u001b[31m"
	green   = "\u001b[32m"
	blue    = "\u001b[34m"
	yellow  = "\u001b[33m"
	megenta = "\u001b[33m"
	end     = "\u001b[0m"
)

var (
	fileData string

	// Errors
	FileError = errors.New(fmt.Sprintf("%sFATT%s: unable to open file", red, end))
)

func readFile(file string, queue chan<- string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(FileError)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		queue <- scanner.Text()
	}
	close(queue)
}

func search(queue <-chan string, finished chan bool) {
	for line := range queue {
		if line == "" {
			continue
		} else {
			for patternName, pattern := range strings.Patterns {
				match, _ := regexp.MatchString(pattern, line)
				if match {
					fmt.Println(patternName, match, line)
				}
			}
		}
	}
	finished <- true
}

func main() {
	c := models.Args{}
	arg.MustParse(&c)

	workQueue := make(chan string)
	finished := make(chan bool)

	if c.File != "" {
		go readFile(c.File, workQueue)
	}

	for i := 0; i < c.Workers; i++ {
		go search(workQueue, finished)
	}

	for i := 0; i < c.Workers; i++ {
		<-finished
	}

	close(finished)
}
