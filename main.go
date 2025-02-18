package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task struct
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// File path for saving JSON data
const filePath = "tasks.json"

// In-memory task list
var tasks []Task

// Load tasks from file
func loadTasks() {
	file, err := ioutil.ReadFile(filePath)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

// Save tasks to file
func saveTasks() {
	file, _ := json.MarshalIndent(tasks, "", "  ")
	ioutil.WriteFile(filePath, file, 0644)
}

// Get all tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// Create a new task
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new ID
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	saveTasks() // Save to file

	c.JSON(http.StatusCreated, newTask)
}

// Update a task
func updateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, task := range tasks {
		if task.ID == id {
			if err := c.ShouldBindJSON(&tasks[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			saveTasks() // Save to file
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Delete a task
func deleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks() // Save to file
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func main() {
	// Load tasks when server starts
	loadTasks()

	r := gin.Default()

	r.GET("/tasks", getTasks)        // Get all tasks
	r.POST("/tasks", createTask)     // Create new task
	r.PUT("/tasks/:id", updateTask)  // Update a task
	r.DELETE("/tasks/:id", deleteTask) // Delete a task

	r.Run(":8080") // Run server on port 8080
}
