package main

import (
	"context"
	"fmt"
)

var cmdList = &Command{
	UsageLine: "list",
	Short:     "list all of the available .gitignore templates",
	Long: `
Lists all of the available .gitignore templates from
https://github.com/github/gitignore.

Example usage:

  gitignore list

`,
}

func init() {
	cmdList.Run = runList
}

func runList(cmd *Command, args []string) {
	ctx := context.Background()
	templates, _, err := client.Gitignores.List(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range templates {
		fmt.Println(t)
	}
}
