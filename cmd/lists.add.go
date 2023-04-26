package cmd

import "github.com/spf13/cobra"

var cmdListsAdd = &cobra.Command{
	Use:   "add",
	Short: "add new list",
	Args:  cobra.MinimumNArgs(1),
}
