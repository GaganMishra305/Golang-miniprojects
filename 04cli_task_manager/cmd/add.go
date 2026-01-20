package cmd

import (
	"fmt"
	"os"
	"strings"

	"task/db"

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

		// Connect to database
		database, err := db.Connect()
		if err != nil {
			fmt.Println("Error connecting to database:", err)
			os.Exit(1)
		}
		defer database.Close()

		// Join all args as the task description
		description := strings.Join(args, " ")

		// Add task to database
		err = db.AddTask(database, description)
		if err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}

		fmt.Printf("Added task: %s\n", description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
