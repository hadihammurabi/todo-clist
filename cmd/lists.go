package cmd

import (
	"fmt"
	"todo-clist/pkg"
	"todo-clist/repo"

	"github.com/spf13/cobra"
)

type CommandLists struct {
	repo    *repo.Repo
	Command *cobra.Command
}

func NewCommandLists(r *repo.Repo) *CommandLists {
	c := &CommandLists{
		r,
		&cobra.Command{
			Use:   "lists",
			Short: "manage lists",
			Args:  cobra.MinimumNArgs(1),
		},
	}

	cmdListsLs.Run = c.GetAllList
	c.Command.Run = c.GetAllList
	c.Command.AddCommand(cmdListsLs)

	cmdListsAdd.Run = c.AddNewList
	c.Command.AddCommand(cmdListsAdd)

	return c
}

func (c *CommandLists) GetAllList(cc *cobra.Command, args []string) {
	lists, err := c.repo.List.GetAll()
	if err != nil {
		panic(fmt.Sprintf("can't read lists: %s", err))
	}

	t := pkg.NewTable("ID", "Name")
	for _, v := range lists {
		t.AddRow(v.ID, v.Name)
	}
	t.Print()
}

func (c *CommandLists) AddNewList(cc *cobra.Command, args []string) {
	err := c.repo.List.Insert(&repo.ListModel{
		Name: args[0],
	})
	if err != nil {
		panic(fmt.Sprintf("can't read lists: %s", err))
	}

	fmt.Println("success: new list created")
}
