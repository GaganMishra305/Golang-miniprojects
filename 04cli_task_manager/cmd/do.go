package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task/db"

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

		// Parse task ID
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a number.")
			os.Exit(1)
		}

		// Connect to database
		database, err := db.Connect()
		if err != nil {
			fmt.Println("Error connecting to database:", err)
			os.Exit(1)
		}
		defer database.Close()

		// Complete the task
		err = db.CompleteTask(database, taskID)
		if err != nil {
			fmt.Println("Error completing task:", err)
			os.Exit(1)
		}

		fmt.Printf("Completed task %d!\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
