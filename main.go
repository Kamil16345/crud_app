package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	//p := prompt.New(
	//	executor,
	//	func(d prompt.Document) []prompt.Suggest { return []prompt.Suggest{} },
	//	prompt.OptionPrefix("CRUD-app"),
	//)
	//p.Run()
	executor()
}

// func executor(s string) {
func executor() {
	fmt.Println("Welcome in CRUD application")
	fmt.Println("---------------------------")
	fmt.Println("What you'd like to do? Choose corresponding number:")
	fmt.Println("1 - Maintain users")
	fmt.Println("2 - Maintain tasks")
	fmt.Println("3 - Quit")
	for {
		operation, err := chooseOperation(`^[1-3]$`)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		operationInt, err := strconv.Atoi(operation)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		switch operationInt {
		case 1:
			for {
				maintainUsers()
			}
		case 2:
			for {
				maintainTasks()
			}
		case 3:
			fmt.Println("Bye!")
			break
		}
	}
}
func chooseOperation(regexExpression string) (string, error) {
	var operation string
	fmt.Scanln(&operation)
	regex := regexp.MustCompile(regexExpression)
	if regex.MatchString(operation) {
		return operation, nil
	}
	return "", fmt.Errorf("Invalid operation, try again.")
}
