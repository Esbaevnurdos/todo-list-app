const API_URL = "https://todo-list-app-hvv4.onrender.com/tasks";

document.addEventListener("DOMContentLoaded", fetchTasks);

const taskForm = document.getElementById("taskForm");
const taskInput = document.getElementById("taskInput");
const taskList = document.getElementById("taskList");

// Fetch all tasks
function fetchTasks() {
  fetch(API_URL)
    .then((response) => response.json())
    .then((tasks) => {
      taskList.innerHTML = "";
      tasks.forEach((task) => addTaskToDOM(task));
    })
    .catch((error) => console.error("Error fetching tasks:", error));
}

// Add new task
taskForm.addEventListener("submit", function (e) {
  e.preventDefault();
  const title = taskInput.value.trim();
  if (title === "") return;

  fetch(API_URL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title: title, done: false }),
  })
    .then((response) => response.json())
    .then((task) => {
      addTaskToDOM(task);
      taskInput.value = "";
    })
    .catch((error) => console.error("Error adding task:", error));
});

// Add task to the DOM
function addTaskToDOM(task) {
  const li = document.createElement("li");
  li.dataset.id = task.id;
  li.className = task.done ? "completed" : "";
  li.innerHTML = `
        <span onclick="toggleTask(${task.id})">${task.title}</span>
        <button class="delete-btn" onclick="deleteTask(${task.id})">Delete</button>
    `;
  taskList.appendChild(li);
}

// Toggle task completion
function toggleTask(id) {
  fetch(`${API_URL}/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ done: true }), // Mark as completed
  })
    .then((response) => response.json())
    .then(() => fetchTasks()) // Refresh list
    .catch((error) => console.error("Error updating task:", error));
}

// Delete task
function deleteTask(id) {
  fetch(`${API_URL}/${id}`, { method: "DELETE" })
    .then(() => fetchTasks()) // Refresh list
    .catch((error) => console.error("Error deleting task:", error));
}
