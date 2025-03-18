---
marp: true
theme: default
class: lead
paginate: true
---

<!-- headingDivider: 1 -->
<!-- backgroundColor: black -->
<!-- class: invert -->

# Warmup

- Phones in bags, bags against the wall
- Create the following and open it in VS Code

```text
/python
  /unit_5
    /lesson_6
```

# Agenda  
- Building CLI apps with OOP and SQLite3  

# Housekeeping  
- Permission forms  

## 🏆 Goal:  
Create a **To-Do List App** using OOP and SQLite3 where the user can:  
- Add a task  
- View all tasks  
- Mark a task as complete  
- Exit the program  

# Step 1: Create the `Task` Class  
Create a `Task` class to store the task details and status:  

```python
class Task:
    def __init__(self, id, description, completed=False):
        self.id = id
        self.description = description
        self.completed = completed
    
    def mark_complete(self):
        self.completed = True
    
    def display(self):
        status = "✓" if self.completed else "✗"
        return f"[{status}] {self.description}"
```

# ✅ **Explanation:**  
- `__init__` sets up the task ID, description, and status  
- `mark_complete()` sets `completed` to `True`  
- `display()` formats the task for display  

# Step 2: Create a `ToDoList` Class  
Create a `ToDoList` class to manage a list of tasks using SQLite3:  

```python
import sqlite3

class ToDoList:
    def __init__(self):
        self.conn = sqlite3.connect('tasks.db')
        self.cursor = self.conn.cursor()
        self.cursor.execute(
            '''CREATE TABLE IF NOT EXISTS tasks (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                description TEXT,
                completed INTEGER
            )'''
        )
        self.conn.commit()
```

# ✅ **Explanation:**  
- `sqlite3.connect()` creates or opens the `tasks.db` file  
- `CREATE TABLE IF NOT EXISTS` ensures the table exists  

# Step 2: Create a `ToDoList` Class (continued)  
Add methods for adding and displaying tasks:  

```python
    def add_task(self, description):
        self.cursor.execute(
            'INSERT INTO tasks (description, completed) VALUES (?, ?)', 
            (description, 0)
        )
        self.conn.commit()
        print(f"Task '{description}' added.")
    
    def display_tasks(self):
        self.cursor.execute('SELECT id, description, completed FROM tasks')
        tasks = self.cursor.fetchall()
        if not tasks:
            print("\nNo tasks available.\n")
            return
        print("\nTasks:")
        for id, description, completed in tasks:
            status = "✓" if completed else "✗"
            print(f"{id}. [{status}] {description}")
```

# ✅ **Explanation:**  
- `INSERT INTO` adds a new task to the database  
- `SELECT` retrieves all tasks from the database  

# Step 2: Create a `ToDoList` Class (continued)  
Add a method to mark a task as complete:  

```python
    def mark_task_complete(self, id):
        self.cursor.execute(
            'UPDATE tasks SET completed = 1 WHERE id = ?', 
            (id,)
        )
        if self.cursor.rowcount:
            print(f"Task '{id}' marked as complete.")
        else:
            print("Invalid task number.")
        self.conn.commit()
```

# ✅ **Explanation:**  
- `UPDATE` modifies the task status  
- `rowcount` checks if the task exists  

# Step 3: Build the Command-Line Interface  
Create a loop to allow the user to interact with the app:  

```python
def main():
    todo = ToDoList()

    while True:
        print("\nOptions:")
        print("1. Add task")
        print("2. View tasks")
        print("3. Mark task complete")
        print("4. Exit")

        choice = input("Enter your choice: ")
```

# Step 3: Build the Command-Line Interface (continued)  
```python
        if choice == "1":
            description = input("Enter task description: ")
            todo.add_task(description)
        elif choice == "2":
            todo.display_tasks()
        elif choice == "3":
            try:
                id = int(input("Enter task number to mark as complete: "))
                todo.mark_task_complete(id)
            except ValueError:
                print("Please enter a valid number.")
        elif choice == "4":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Try again.")
```

# Step 3: Build the Command-Line Interface (continued)  
```python
if __name__ == "__main__":
    main()
```

# ✅ **Explanation:**  
- The loop keeps running until the user exits  
- User input is handled by `ToDoList` methods  
- `try/except` handles invalid inputs  

# ✅ Sample Run  
```text
Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 1
Enter task description: Finish homework
Task 'Finish homework' added.

Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 2

Tasks:
1. [✗] Finish homework

Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 3
Enter task number to mark as complete: 1
Task '1' marked as complete.

Options:
1. Add task
2. View tasks
3. Mark task complete
4. Exit
Enter your choice: 4
Goodbye!
```

# ✅ Done!  
✅ Fully working app using OOP and SQLite3  
✅ Handles edge cases like invalid inputs  
✅ Clean and readable code  
