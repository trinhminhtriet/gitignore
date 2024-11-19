package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/google/go-github/github"
)

var cmdCreate = &Command{
	UsageLine: "create [languages]",
	Short:     "create a .gitignore file",
	Long: `
Create downloads the .gitignore templates for the languages specified from
https://github.com/github/gitignore and mashes them together into a .gitignore
file.

Example usage:

  gitignore create go

`,
}

func init() {
	cmdCreate.Run = runCreate
}

func runCreate(cmd *Command, args []string) {
	if _, err := os.Stat(".gitignore"); err == nil {
		fmt.Println("Error: A .gitignore file already exists.")
		os.Exit(2)
	}

	ctx := context.Background()
	availableLanguages, _, err := client.Gitignores.List(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}

	chosenLanguages := make([]*string, len(args))
OUTER:
	for i, n := range args {
		for _, a := range availableLanguages {
			if strings.EqualFold(n, a) {
				chosenLanguages[i] = &a
				continue OUTER
			}
		}
		fmt.Fprintf(os.Stderr, "%s is not a valid template!\n", n)
		os.Exit(2)
	}

	var wg sync.WaitGroup

	gitignores := make([]*github.Gitignore, len(chosenLanguages))
	for index, name := range chosenLanguages {
		wg.Add(1)

		go func(n string, i int, gitignores []*github.Gitignore) {
			defer wg.Done()
			fmt.Println("Fetching template for " + n + "...")
			ctx := context.Background()
			gitignore, _, err := client.Gitignores.Get(ctx, n)
			if err != nil {
				fmt.Println(err)
				return
			}
			gitignores[i] = gitignore
		}(*name, index, gitignores)
	}

	wg.Wait()

	var fileContents string
	for _, g := range gitignores {
		header := header(*g.Name)
		fileContents += header + "\n\n" + *g.Source + "\n\n"
	}

	fileContents = strings.TrimSpace(fileContents)

	fmt.Println("Writing .gitignore file...")
	err = ioutil.WriteFile(".gitignore", []byte(fileContents), 0644)
	if err != nil {
		fmt.Println("Error writing .gitignore file!")
		os.Exit(2)
	}
	fmt.Println("Done!")
}

func header(name string) string {
	line := strings.Repeat("#", len(name)+8)
	mid := "#   " + name + "   #"
	return line + "\n" + mid + "\n" + line
}
