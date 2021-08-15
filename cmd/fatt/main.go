package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/binexisHATT/fatt/helpers"
	"github.com/binexisHATT/fatt/models"
	"github.com/binexisHATT/fatt/strings"
	"github.com/dlclark/regexp2"
)

const (
	red     = "\u001b[31m"
	green   = "\u001b[32m"
	blue    = "\u001b[34m"
	yellow  = "\u001b[33m"
	megenta = "\u001b[33m"
	end     = "\u001b[0m"
	bold    = "\u001b[1m"
	underL  = "\u001b[4m"
)

var (
	fileData string
	noColor  bool

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
				re := regexp2.MustCompile(pattern, 0)
				matches := helpers.FindAllString(re, line)
				for _, match := range matches {
					if noColor {
						log.Println(
							fmt.Sprintf("%s:%s", patternName, match),
						)
					} else {
						log.Println(
							fmt.Sprintf("%s%s%s%s:%s", green, underL, patternName, end, match),
						)
					}
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

	if c.NoColor {
		noColor = true
	}

	if noColor {
		log.SetPrefix("fatt:")
		log.SetFlags(0)
	} else {
		log.SetPrefix(fmt.Sprintf("%s%sfatt%s:", red, bold, end))
		log.SetFlags(0)
	}

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
