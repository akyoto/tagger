package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	tags         bool
	root         string
	shellCommand string
)

func init() {
	flag.BoolVar(&tags, "tags", false, "Shows the latest tag in every git repository")
	flag.StringVar(&root, "root", ".", "Specifies the directory to search for git repositories")
	flag.StringVar(&shellCommand, "run", "", "Specifies a shell command to execute in every git repository")
	flag.Parse()
}

func main() {
	// Show help
	if !tags && shellCommand == "" {
		flag.Usage()
		return
	}

	// Process repositories in parallel
	err := filepath.Walk(root, walk)

	if err != nil {
		panic(err)
	}

	repositoryWaitGroup.Wait()

	// Shell command output
	if shellCommand != "" {
		fmt.Print("\033[2K\r")
		showCommandOutput()
	}

	// Tags output
	if tags {
		if shellCommand != "" {
			fmt.Println()
		}

		showTags()
	}
}
