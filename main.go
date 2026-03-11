package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func main() {
	journalFile, err := getJournalFilePath()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: could not find Journal file:", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "-p" || os.Args[1] == "--post" {
			postStr := strings.Join(os.Args[2:], " ")
			appendPost(journalFile, postStr)
			os.Exit(0)
		}
	}

	searchPattern := buildSearchPattern()
	displayJournal(journalFile, searchPattern)
}

func displayJournal(journalFile string, searchPattern *regexp.Regexp) {
	file, err := os.OpenFile(journalFile, os.O_RDONLY, fs.FileMode(os.O_SYNC))
	if err != nil {
		fmt.Printf("error: could not read file %s\n", journalFile)
		os.Exit(2)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		if searchPattern.Match(line) {
			fmt.Printf("\033[0;94m%s\033[0m%s\n", line[0:16], line[16:])
		}
	}

	file.Close()
}

func buildSearchPattern() *regexp.Regexp {
	combinedArgs := strings.Join(os.Args[1:], "|")
	if combinedArgs == "today" {
		combinedArgs = time.Now().Format(time.DateOnly)
	} else if combinedArgs == "" {
		combinedArgs = ".*"
	}

	searchPattern, err := regexp.Compile(combinedArgs)
	if err != nil {
		fmt.Printf("error: could not build search expression from '%s'\n", combinedArgs)
		os.Exit(3)
	}

	return searchPattern
}

func appendPost(journalFile string, post string) {
	file, err := os.OpenFile(journalFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("error: could not read file %s\n", journalFile)
		os.Exit(2)
	}

	datestamp := time.Now().Format("2006-01-02 15:04")

	_, err = file.Write([]byte(datestamp + " " + post + "\n"))
	if err != nil {
		fmt.Printf("error: could not write to file %s\n", journalFile)
		os.Exit(3)
	}

	file.Close()
}

func getXdgDataHome() string {
	xdgDataHome, ok := os.LookupEnv("XDG_DATA_HOME")
	if !ok {
		homeDir, _ := os.LookupEnv("HOME")
		xdgDataHome = path.Join(homeDir, ".local/share")
	}

	return xdgDataHome
}

func getJournalFilePath() (string, error) {
	xdgDataHome := getXdgDataHome()
	journalDataDir := path.Join(xdgDataHome, "journal")

	if err := os.MkdirAll(journalDataDir, fs.ModePerm); err != nil {
		return "", err
	}

	return path.Join(journalDataDir, "data.journal"), nil
}
