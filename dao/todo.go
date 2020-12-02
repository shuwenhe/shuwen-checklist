package dao

import (
	db "github.com/shuwenhe/shuwen-checklist/database"
	"github.com/shuwenhe/shuwen-checklist/model"
)

func CreateTodo(todo *model.Todo) (err error) { // todo C
	err = db.DB.Create(&todo).Error
	return
}

func GetTodoList() (todoList []*model.Todo, err error) {
	if err = db.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return // if successful data return
}

func GetTodoById(id string) (todo *model.Todo, err error) {
	todo = new(model.Todo)                                                     // It will report an error if it is not instantiated by the new keyword
	if err := db.DB.Debug().Where("id=?", id).First(&todo).Error; err != nil { // 到数据库查
		return nil, err
	}
	return
}

func UpdateTodo(todo *model.Todo) (err error) { // Returns whether the update succeeded or failed
	err = db.DB.Save(&todo).Error
	return
}
func DeleteTodoById(id string) (err error) {
	err = db.DB.Where("id = ?", id).Delete(&model.Todo{}).Error
	return
}
