package main

import (
	"fmt"
	"strconv"
)

func maintainTasks() {
	for {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("What would you like to do? Choose corresponding number:")
		fmt.Println("1 - Create task")
		fmt.Println("2 - Display particular user tasks")
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

			fmt.Println("Type in user ID (required field): ")
			_, err = fmt.Scanln(&task.User_Id)

			fmt.Println("Type in task title: ")
			_, err = fmt.Scanln(&task.Title)

			fmt.Println("Type in task description: ")
			_, err = fmt.Scanln(&task.Description)

			err := task.Create(task)

			if err != nil {
				fmt.Println("Error creating task: ", err)
			} else {
				fmt.Println("Task created successfully!")
			}
		case 2:
			task := &Task{}
			fmt.Println("Type in user ID to display its tasks: ")
			_, err = fmt.Scanln(&task.Id)
			_, err := task.Read(task.Id)
			if err != nil {
				fmt.Println("Error reading task: ", err)
			}
		case 3:
			task := &Task{}
			_, err := task.ReadAll(true)
			if err != nil {
				fmt.Println("Error reading task: ", err)
			}
		case 4:
			task := &Task{}

			fmt.Println("Type in task id to update: ")
			_, err = fmt.Scanln(&task.Id)
			exists := checkIfTaskExists(task.Id)
			if exists {
				fmt.Println("New assigned user (field required): ")
				_, err = fmt.Scanln(&task.User_Id)
				userExists := checkIfUserExists(task.User_Id)
				if userExists {
					fmt.Println("New title(leave blank, if unnecessary): ")
					_, err = fmt.Scanln(&task.Title)
					fmt.Println("New description(leave blank, if unnecessary): ")
					_, err = fmt.Scanln(&task.Description)
				}

				err := task.Update(task.Id, task)
				if err != nil {
					fmt.Println("Error updating task: ", err)
				} else {
					fmt.Println("Updating a task went successfully.")
				}
			} else {
				fmt.Println("Task with given ID not found")
			}
		case 5:
			task := &Task{}
			fmt.Println("Type in task name to delete: ")
			_, err = fmt.Scanln(&task.Id)
			err := task.Delete(task.Id)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
			} else {
				fmt.Println("Removing a task went successfully.")
			}
		case 6:
			executor()
		}
	}
}
func maintainUsers() {
	for {
		fmt.Println("--------------------------------------------------------")
		fmt.Println("What would you like to do? Choose corresponding number: ")
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

			fmt.Println("Type in user name: ")
			_, err = fmt.Scanln(&user.Name)
			if err != nil {
				fmt.Println("Error reading user name: ", err)
				return
			}

			fmt.Println("Type in user email: ")
			_, err = fmt.Scanln(&user.Email)
			if err != nil {
				fmt.Println("Error reading user email: ", err)
				return
			}

			fmt.Println("Type in user phone: ")
			_, err = fmt.Scanln(&user.Phone)
			if err != nil {
				fmt.Println("Error reading user phone: ", err)
			}
			fmt.Println("user in operations.go", user)
			err := user.Create(user)

			if err != nil {
				fmt.Println("Error creating user: ", err)
			} else {
				fmt.Println("User created successfully!")
			}
		case 2:
			user := &User{}
			fmt.Println("Type in user id to display: ")
			_, err = fmt.Scanln(&user.Id)
			_, err := user.Read(user.Id)
			if err != nil {
				fmt.Println("Error reading user: ", err)
			} else {
				fmt.Println("User found")
			}
		case 3:
			user := &User{}
			_, err := user.ReadAll(true)
			if err != nil {
				fmt.Println("Error reading users: ", err)
			}
		case 4:
			user := &User{}

			fmt.Println("Type in user id to update: ")
			_, err = fmt.Scanln(&user.Id)
			exists := checkIfUserExists(user.Id)
			if exists {
				fmt.Println("New assigned user name(leave blank, if unnecessary): ")
				_, err = fmt.Scanln(&user.Name)

				fmt.Println("New email(leave blank, if unnecessary): ")
				_, err = fmt.Scanln(&user.Email)

				fmt.Println("New phone(leave blank, if unnecessary): ")
				_, err = fmt.Scanln(&user.Phone)

				err := user.Update(user.Id, user)
				if err != nil {
					fmt.Println("Error updating task: ", err)
				} else {
					fmt.Println("Updating a task went successfully.")
				}
			} else {
				fmt.Println("User with given ID not found")
			}

		case 5:
			user := &User{}
			fmt.Println("Type in user id to delete: ")
			_, err = fmt.Scanln(&user.Id)
			err := user.Delete(user.Id)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
			} else {
				fmt.Println("Removing a task went successfully.")
			}
		case 6:
			executor()
		}
	}
}
