package cmd

import (
	"todo-clist/repo"

	"github.com/spf13/cobra"
)

type CommandTodos struct {
	repo    *repo.Repo
	Command *cobra.Command
}

func NewCommandTodos(r *repo.Repo) *CommandTodos {
	return &CommandTodos{
		r,
		&cobra.Command{
			Use:   "todo",
			Short: "manage todos",
			Run: func(c *cobra.Command, args []string) {
				println("manage all todos")
			},
		},
	}
}
