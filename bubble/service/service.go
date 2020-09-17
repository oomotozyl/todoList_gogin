package service

import (
	"bubble/models"
)

// SelectAll SelectAll
func SelectAll() (todoList []*models.Todo, err error) {
	return models.SelectAll()
}

// CreateOne CreateOne
func CreateOne(todo *models.Todo) error {
	return models.CreateOne(todo)
}

// SelectOne SelectOne
func SelectOne(id string) (todo *models.Todo, err error) {
	return models.SelectOne(id)
}

// UpdateOne UpdateOne
func UpdateOne(todo *models.Todo) error {
	return models.UpdateOne(todo)
}

// DeleteOne DeleteOne
func DeleteOne(id string) error {
	return models.DeleteOne(id)
}
