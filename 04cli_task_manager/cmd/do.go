package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Doing a specific task in the queue.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Give one task id to do. [Ex: task do 1].")
			os.Exit(1)
		}
		fmt.Println("do called")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
