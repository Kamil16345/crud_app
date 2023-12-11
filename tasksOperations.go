package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Task struct {
	id          int    `json:"id"`
	user_id     int    `json:"user_Id"`
	title       string `json:"title"`
	description string `json:"description"`
}

func (t *Task) Create(data interface{}) error {
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer file.Close()
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *Task) Read(id int) (interface{}, error) {
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return nil, err
	}
	task := findTask(tasks, id)
	if task != nil {
		fmt.Println("Task found:")
		fmt.Println("Task ID:", task.id)
		fmt.Println("Assigned user ID:", task.user_id)
		fmt.Println("Title:", task.title)
		fmt.Println("Description:", task.description)
		return task, nil
	} else {
		fmt.Println("Task not found")
		return nil, err
	}

}
func (t *Task) ReadAll() (interface{}, error) {
	filePath := "tasks.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error during unmarshalling JSON: ", err)
		return nil, err
	}
	for _, task := range tasks {
		fmt.Println("Task ID:", task.id)
		fmt.Println("Assigned user ID:", task.user_id)
		fmt.Println("Title:", task.title)
		fmt.Println("Description:", task.description)
		fmt.Println("-------------------------------")
	}
	return tasks, err
}
func (t *Task) Update(id int, data interface{}) error {
	filePath := "tasks.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return err
	}
	var tasks []Task
	dataBytes := data.([]byte)
	err = json.Unmarshal(dataBytes, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return err
	}
	newTask := data.(*Task)
	task := findTask(tasks, id)
	if task == nil {
		fmt.Println("Task not found")
		return nil
	}
	task.user_id = newTask.user_id
	task.title = newTask.title
	task.description = newTask.description

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshalling JSON data: ", err)
		return err
	}
	err = ioutil.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *Task) Delete(id int) error {
	data, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error in unmarshalling JSON data: ", err)
		return err
	}
	task := findTask(tasks, id)
	if task == nil {
		fmt.Println("Task not found")
		return nil
	}
	for i, task := range tasks {
		if task.id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshalling JSON data: ", err)
		return err
	}

	err = ioutil.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}

	fmt.Println("Task deleted successfully")
	return nil
}

func findTask(tasks []Task, id int) *Task {
	for _, task := range tasks {
		if task.id == id {
			return &task
		}
	}
	return nil
}
