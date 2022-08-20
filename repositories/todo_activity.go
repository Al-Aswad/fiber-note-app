package repositories

import (
	"al-aswad/fiber-note-app/models"
	"log"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Update(id int, todo models.Todo) (models.Todo, bool)
	Delete(id int) (bool, interface{})
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) Create(todo models.Todo) (models.Todo, error) {
	err := t.db.Save(&todo).Error

	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil

}

func (t *todoRepository) GetAll() ([]models.Todo, error) {
	todo := []models.Todo{}

	err := t.db.Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *todoRepository) Update(id int, todo models.Todo) (models.Todo, bool) {
	var todoUpdate models.Todo

	err := t.db.Debug().Model(&todoUpdate).Where("id = ?", id).Updates(&todo)
	if err.RowsAffected == 0 {
		log.Println("Update Activity ", err.Error)
		return models.Todo{}, false
	}

	todoUpdate.ID = uint(id)
	return todoUpdate, true
}

func (t *todoRepository) Delete(id int) (bool, interface{}) {
	errFound := t.db.Delete(&models.Todo{}, id)

	if errFound.RowsAffected == 0 {
		return false, errFound.Error
	}

	return true, nil
}
