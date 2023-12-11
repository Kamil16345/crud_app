package main

import (
	"fmt"
	"strconv"
)

func maintainTasks() {
	for {
		fmt.Println("What would you like to do? Choose corresponding number:")
		fmt.Println("1 - Create task")
		fmt.Println("2 - Display particular task")
		fmt.Println("3 - Display tasks")
		fmt.Println("4 - Update task")
		fmt.Println("5 - Delete task")
		fmt.Println("6 - Quit")
		operation, err := chooseOperation(`^[1-6]$`)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		operationInt, err := strconv.Atoi(operation)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		switch operationInt {
		case 1:
			task := &Task{}

			fmt.Println("Type in task id: ")
			fmt.Scanf("%s", &task.id)

			fmt.Println("Type in user id: ")
			fmt.Scanf("%s", &task.user_id)

			fmt.Println("Type in task title: ")
			fmt.Scanf("%s", &task.title)

			fmt.Println("Type in task description: ")
			fmt.Scanf("%s", &task.description)

			err := task.Create(&task)

			if err != nil {
				fmt.Println("Error creating task: ", err)
			} else {
				fmt.Println("Task created successfully!")
			}
		case 2:
			task := &Task{}

			fmt.Println("Type in task title to display: ")
			fmt.Scanf("%d", &task.id)

			data, err := task.Read(task.id)
			if err != nil {
				fmt.Println("Error reading task: ", err)
			} else {
				fmt.Println("Here is a task: ", data)
			}
		case 3:
			task := &Task{}
			data, err := task.ReadAll()
			if err != nil {
				fmt.Println("Error reading task: ", err)
			} else {
				fmt.Println("Here are all tasks: ", data)
			}
		case 4:
			task := &Task{}

			fmt.Println("Type in task id to update: ")
			fmt.Scanf("%d", &task.id)

			fmt.Println("New assigned user(leave blank, if unnecessary): ")
			fmt.Scanf("%d", &task.user_id)

			fmt.Println("New title(leave blank, if unnecessary): ")
			fmt.Scanf("%s", &task.title)

			fmt.Println("New description(leave blank, if unnecessary): ")
			fmt.Scanf("%s", &task.description)

			err := task.Update(task.id, task)
			if err != nil {
				fmt.Println("Error updating task: ", err)
			} else {
				fmt.Println("Updating a task went successfully.")
			}
		case 5:
			task := &Task{}
			fmt.Println("Type in task name to delete: ")
			fmt.Scanf("%d", &task.id)
			err := task.Delete(task.id)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
			} else {
				fmt.Println("Removing a task went successfully.")
			}
		case 6:
			break
		}
	}
}
func maintainUsers() {
	for {
		fmt.Println("What would you like to do? Choose corresponding number:")
		fmt.Println("1 - Create user")
		fmt.Println("2 - Display particular user")
		fmt.Println("3 - Display users")
		fmt.Println("4 - Update user")
		fmt.Println("5 - Delete user")
		fmt.Println("6 - Quit")
		operation, err := chooseOperation(`^[1-6]$`)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		operationInt, err := strconv.Atoi(operation)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		switch operationInt {
		case 1:
			user := &User{}

			fmt.Println("Type in user id: ")
			user.id, err = fmt.Scanln(&user.id)

			fmt.Println("Type in user name: ")
			name, err := fmt.Scanln(&user.name)
			if err != nil {
				fmt.Println("Error reading user name: ", err)
				return
			}

			fmt.Println("Type in user email: ")
			user.email, err := fmt.Scanln(&user.email)
			if err != nil {
				fmt.Println("Error reading user email: ", err)
				return
			}

			fmt.Println("Type in user phone: ")
			phone, err := fmt.Scanln(&user.phone)
			if err != nil {
				fmt.Println("Error reading user phone: ", err)
			}

			err := user.Create(&user)

			if err != nil {
				fmt.Println("Error creating user: ", err)
			} else {
				fmt.Println("User created successfully!")
			}
		case 2:
			user := &User{}

			fmt.Println("Type in user id to display: ")
			fmt.Scanf("%d", &user.id)

			data, err := user.Read(user.id)
			if err != nil {
				fmt.Println("Error reading user: ", err)
			} else {
				fmt.Println("Here is a user: ", data)
			}
		case 3:
			user := &User{}
			data, err := user.ReadAll()
			if err != nil {
				fmt.Println("Error reading users: ", err)
			} else {
				fmt.Println("Here are all users: ", data)
			}
		case 4:
			user := &User{}

			fmt.Println("Type in user id to update: ")
			fmt.Scanf("%d", &user.id)

			fmt.Println("New assigned user name(leave blank, if unnecessary): ")
			fmt.Scanf("%d", &user.name)

			fmt.Println("New email(leave blank, if unnecessary): ")
			fmt.Scanf("%s", &user.email)

			fmt.Println("New phone(leave blank, if unnecessary): ")
			fmt.Scanf("%s", &user.phone)

			err := user.Update(user.id, user)
			if err != nil {
				fmt.Println("Error updating task: ", err)
			} else {
				fmt.Println("Updating a task went successfully.")
			}
		case 5:
			user := &User{}
			fmt.Println("Type in user id to delete: ")
			fmt.Scanf("%d", &user.id)
			err := user.Delete(user.id)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
			} else {
				fmt.Println("Removing a task went successfully.")
			}
		case 6:
			break
		}
	}
}
