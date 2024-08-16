package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
  var newPost = flag.String("p", "", "Post content")
  flag.Parse()

	journalFile, ok := os.LookupEnv("JOURNAL_DATA_FILE")
	if !ok {
		fmt.Printf("error: could not find JOURNAL_DATA_FILE in env")
		os.Exit(1)
	}

  if flag.NFlag() > 0 {
    appendPost(journalFile, *newPost)
    os.Exit(0)
  }

  searchPattern := buildSearchPattern()
  displayJournal(journalFile, searchPattern)
}

func displayJournal(journalFile string, searchPattern *regexp.Regexp) {
  file, err := os.OpenFile(journalFile, os.O_RDONLY, fs.FileMode(os.O_SYNC)) 
	if err != nil {
		fmt.Printf("error: could not read file %s", journalFile)
		os.Exit(2)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		if searchPattern.Match(line) {
			fmt.Println(string(line))
		}
	}

	file.Close()
}

func buildSearchPattern() *regexp.Regexp {
	combinedArgs := strings.Join(os.Args[1:], "|")
	if combinedArgs == "" {
		combinedArgs = time.Now().Format(time.DateOnly)
	}

	searchPattern, err := regexp.Compile(combinedArgs)
	if err != nil {
		fmt.Printf("error: could not build search expression from '%s'", combinedArgs)
		os.Exit(3)
	}

  return searchPattern
}

func appendPost(journalFile string, post string) {
    file, err := os.OpenFile(journalFile, os.O_APPEND | os.O_WRONLY | os.O_CREATE, fs.FileMode(os.O_RDWR))
    if err != nil {
      fmt.Printf("error: could not read file %s", journalFile)
      os.Exit(2)
    }

    datestamp := time.Now().Format("2006-01-02 15:04")

    _, err = file.Write([]byte(datestamp + " " + post + "\n"))
    if err != nil {
      fmt.Printf("error: could not write to file %s", journalFile)
      os.Exit(3)
    }

    file.Close()
}

