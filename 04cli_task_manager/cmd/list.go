package cmd

import (
	"fmt"
	"os"

	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listing all the tasks in the queue.",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to database
		database, err := db.Connect()
		if err != nil {
			fmt.Println("Error connecting to database:", err)
			os.Exit(1)
		}
		defer database.Close()

		// Get all tasks
		tasks, err := db.ListTasks(database)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks in your queue!")
			return
		}

		fmt.Println("Your tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s\n", task.ID, task.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
