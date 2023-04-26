package main

import (
	"todo-clist/cmd"
	"todo-clist/repo"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type App struct {
	repo *repo.Repo
	cm   *cobra.Command
}

func NewApp(db *gorm.DB) *App {
	r := repo.NewRepo(db)

	defaultCommand := cmd.NewCommandTodos(r)
	app := &cobra.Command{
		Use:   "todo-clist",
		Short: "Manage activities easily",
		Run: func(c *cobra.Command, args []string) {
			defaultCommand.Command.Run(c, append(args, "ls"))
		},
	}
	app.AddCommand(defaultCommand.Command)
	app.AddCommand(cmd.NewCommandLists(r).Command)

	return &App{
		repo.NewRepo(db),
		app,
	}
}

func (a *App) Execute() error {
	return a.cm.Execute()
}
