package models

import "bubble/dao"

// Todo Todo
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// SelectAll SelectAll
func SelectAll() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	} else {
		return todoList, nil
	}
}

// CreateOne CreateOne
func CreateOne(todo *Todo) error {
	return dao.DB.Create(&todo).Error
}

// SelectOne SelectOne
func SelectOne(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.First(todo, id).Error
	if err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

// UpdateOne UpdateOne
func UpdateOne(todo *Todo) error {
	return dao.DB.Save(&todo).Error
}

// DeleteOne DeleteOne
func DeleteOne(id string) error {
	return dao.DB.Delete(&Todo{}, id).Error
}
