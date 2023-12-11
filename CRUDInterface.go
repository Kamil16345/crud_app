package main

type CRUD interface {
	Create(data interface{}) error
	Read(id string) (interface{}, error)
	ReadAll() (interface{}, error)
	Update(id string, data interface{}) error
	Delete(id string) error
}
