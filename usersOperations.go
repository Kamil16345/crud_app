package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
	phone string `json:"phone"`
}

func (t *User) Create(data interface{}) error {
	file, err := os.OpenFile("users.json", os.O_WRONLY|os.O_CREATE, 0644)
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
func (t *User) Read(id int) (interface{}, error) {
	data, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return nil, err
	}
	user := findUser(users, id)
	if user != nil {
		fmt.Println("User found:")
		fmt.Println("User ID:", user.id)
		fmt.Println("User name:", user.name)
		fmt.Println("User email:", user.email)
		fmt.Println("User phone:", user.phone)
		return user, nil
	} else {
		fmt.Println("User not found")
		return nil, err
	}
}
func (t *User) ReadAll() (interface{}, error) {
	filePath := "users.json"
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	var users string
	err = json.Unmarshal(fileContent, &users)
	if err != nil {
		fmt.Println("Error during unmarshalling JSON: ", err)
	}
	return users, err
}
func (t *User) Update(id int, data interface{}) error {
	filePath := "users.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return err
	}
	var users []User
	dataBytes := data.([]byte)
	err = json.Unmarshal(dataBytes, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return err
	}
	newUser := data.(*User)
	user := findUser(users, id)
	if user == nil {
		fmt.Println("User not found")
		return nil
	}
	user.name = newUser.name
	user.email = newUser.email
	user.phone = newUser.phone

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error marshalling JSON data: ", err)
		return err
	}
	err = ioutil.WriteFile("users.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *User) Delete(id int) error {
	data, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error in unmarshalling JSON data: ", err)
		return err
	}
	user := findUser(users, id)
	if user == nil {
		fmt.Println("User not found")
		return nil
	}
	for i, user := range users {
		if user.id == id {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error marshalling JSON data: ", err)
		return err
	}

	err = ioutil.WriteFile("users.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}

	fmt.Println("User deleted successfully")
	return nil
}

func findUser(users []User, id int) *User {
	for _, user := range users {
		if user.id == id {
			return &user
		}
	}
	return nil
}
