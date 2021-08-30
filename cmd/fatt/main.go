package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/binexisHATT/fatt/helpers"
	"github.com/binexisHATT/fatt/models"
	"github.com/binexisHATT/fatt/patterns"
	"github.com/dlclark/regexp2"
)

const version = "v1.1.0"

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	blue   = "\u001b[94m"
	yellow = "\u001b[33m"
	end    = "\u001b[0m"
	bold   = "\u001b[1m"
	underL = "\u001b[4m"
)

var (
	noColor           bool
	totalStringsFound int
	quiet             bool
	excludes          []string

	// Errors
	FileError      = errors.New("unable to open file")
	HTTPGetError   = errors.New("unable to make GET requests to specified url")
	ReadStdInError = errors.New("failed to read input from stdin")
)

func search(data string) {
	for patternName, pattern := range patterns.Patterns {
		if helpers.StringInSlice(strings.ToLower(patternName), excludes) {
			continue
		}
		re := regexp2.MustCompile(pattern, 0)
		matches := helpers.FindAllString(re, data)
		for _, match := range matches {
			totalStringsFound += 1
			match := strings.TrimRight(match, "\n")
			if noColor && quiet {
				fmt.Println(
					fmt.Sprintf("[%s]%s", patternName, match),
				)
			} else if noColor {
				log.Println(
					fmt.Sprintf("[%s]%s", patternName, match),
				)
			} else if quiet {
				fmt.Println(
					fmt.Sprintf("[%s%s%s%s]%s%s%s", underL, blue, patternName, end, yellow, match, end),
				)
			} else {
				log.Println(
					fmt.Sprintf("[%s%s%s%s]%s%s%s", underL, blue, patternName, end, yellow, match, end),
				)
			}
		}
	}
}

func scanUrl(url string, queue chan<- string) {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(HTTPGetError)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		queue <- scanner.Text()
	}
	close(queue)
}

func readFile(file string, queue chan<- string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(FileError)
	}
	fileStat, err := f.Stat()
	if err != nil {
		log.Fatalln("unable to determine if argument to `--file` is a file or directory")
	}
	if fileStat.IsDir() {
		log.Fatalln("argument to `--file` must be a file and not a directory")
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

	if c.ListPatterns {
		helpers.ListPatterns()
		return
	}

	if c.NoColor {
		noColor = true
	}

	if c.Quiet {
		quiet = true
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

	if c.Version {
		log.Printf(" %s\n", version)
		return
	}

	if c.Exclude != "" {

		if helpers.FileExists(c.Exclude) {
			f, err := os.Open(c.Exclude)
			if err != nil {
				log.Printf("unable to read file: %s...%sfailure%s", c.Exclude, red, end)
			}
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				excludes = append(excludes, strings.Trim(scanner.Text(), " \n\t"))
			}
		} else {
			excludeSlice := strings.Split(strings.Trim(c.Exclude, " \n\t"), ",")
			for _, str := range excludeSlice {
				excludes = append(excludes, strings.ToLower(str))
			}
		}
	}

	workQueue := make(chan string)
	finished := make(chan bool)

	if c.Url != "" {
		go scanUrl(c.Url, workQueue)
	} else if c.File != "" {
		go readFile(c.File, workQueue)
	} else {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeCharDevice == 0 {
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
		} else {
			log.Fatalln("--file or data from stdin must be specified")
		}
	}

	for i := 0; i < c.Workers; i++ {
		go run(workQueue, finished)
	}

	for i := 0; i < c.Workers; i++ {
		<-finished
	}

	close(finished)

	if quiet {
		return
	}

	results := fmt.Sprintf("\u2022 Total Strings Found: %d", totalStringsFound)
	elapsedTime := fmt.Sprintf("\u2022 Elapsed Time: %s", time.Since(start))
	fmt.Printf("\n%s\n%s\n", results, elapsedTime)
}
