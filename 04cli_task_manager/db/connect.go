package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

var bucketName = []byte("tasks")

type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// Connect opens and returns a database connection
func Connect() (*bolt.DB, error) {
	db, err := bolt.Open("db/tasks.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	// Create bucket if it doesn't exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// AddTask adds a new task to the database
func AddTask(db *bolt.DB, description string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		// Generate ID
		id, _ := b.NextSequence()

		task := Task{
			ID:          int(id),
			Description: description,
			Completed:   false,
			CreatedAt:   time.Now(),
		}

		// Encode task as JSON
		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}

		// Store task
		return b.Put(itob(int(id)), encoded)
	})
}

// ListTasks returns all incomplete tasks
func ListTasks(db *bolt.DB) ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		return b.ForEach(func(k, v []byte) error {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}

			if !task.Completed {
				tasks = append(tasks, task)
			}
			return nil
		})
	})

	return tasks, err
}

// CompleteTask marks a task as completed
func CompleteTask(db *bolt.DB, id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		// Get the task
		v := b.Get(itob(id))
		if v == nil {
			return fmt.Errorf("task with ID %d not found", id)
		}

		var task Task
		if err := json.Unmarshal(v, &task); err != nil {
			return err
		}

		// Mark as completed
		task.Completed = true

		// Encode and save
		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(itob(id), encoded)
	})
}

// itob converts int to byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
