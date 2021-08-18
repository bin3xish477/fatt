package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/binexisHATT/fatt/helpers"
	"github.com/binexisHATT/fatt/models"
	"github.com/binexisHATT/fatt/patterns"
	"github.com/dlclark/regexp2"
)

const (
	red    = "\u001b[31m"
	green  = "\u001b[32m"
	blue   = "\u001b[34m"
	yellow = "\u001b[33m"
	end    = "\u001b[0m"
	bold   = "\u001b[1m"
	underL = "\u001b[4m"
)

var (
	fileData          string
	noColor           bool
	totalStringsFound int

	// Errors
	FileError      = errors.New(" unable to open file")
	HTTPGetError   = errors.New(" unable to make GET requests to specified url")
	ReadStdInError = errors.New(" failed to read input from stdin")
)

func search(data string) {
	for patternName, pattern := range patterns.Patterns {
		re := regexp2.MustCompile(pattern, 0)
		matches := helpers.FindAllString(re, data)
		for _, match := range matches {
			totalStringsFound += 1
			if noColor {
				log.Println(
					fmt.Sprintf("%s:%s", patternName, match),
				)
			} else {
				log.Println(
					fmt.Sprintf("%s%s%s%s:%s%s%s", underL, green, patternName, end, yellow, match, end),
				)
			}
		}
	}
}

func scanUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(HTTPGetError)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
	}
	search(string(body))
}

func readFile(file string, queue chan<- string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(FileError)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		queue <- scanner.Text()
	}
	close(queue)
}

func run(queue <-chan string, finished chan bool) {
	for line := range queue {
		if line == "" {
			continue
		} else {
			search(line)
		}
	}
	finished <- true
}

func main() {
	start := time.Now()
	c := models.Args{}
	arg.MustParse(&c)

	workQueue := make(chan string)
	finished := make(chan bool)

	if c.NoColor {
		noColor = true
	}

	if c.OutFile != "" {
		logFile, err := os.OpenFile(c.OutFile, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalln(fmt.Sprintf("could not create file -> %s", c.OutFile))
		}
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
	} else {
		log.SetOutput(os.Stdout)
	}

	if noColor {
		log.SetPrefix("fatt:")
		log.SetFlags(0)
	} else {
		log.SetPrefix(fmt.Sprintf("%s%sfatt%s:", red, bold, end))
		log.SetFlags(0)
	}

	if c.Url != "" {
		scanUrl(c.Url)
		return
	}

	if c.File != "" {
		go readFile(c.File, workQueue)
	} else {
		// scanning from stdin
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				workQueue <- scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				log.Fatalln(ReadStdInError)
			}
			close(workQueue)
		}()
	}

	for i := 0; i < c.Workers; i++ {
		go run(workQueue, finished)
	}

	for i := 0; i < c.Workers; i++ {
		<-finished
	}

	close(finished)

	fmt.Println(
		fmt.Sprintf("\nTotal Strings Found: %d\nElapsed Time: %s", totalStringsFound, time.Since(start)),
	)
}
