package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding a task to the queue",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Give atleast one task to add. [Ex: task add wash the car].")
			os.Exit(1)
		}
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
