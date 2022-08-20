package services

import (
	"al-aswad/fiber-note-app/models"
	"al-aswad/fiber-note-app/repositories"
	"al-aswad/fiber-note-app/requests"
	"log"

	"github.com/mashingan/smapping"
)

type TodoService interface {
	Create(todo requests.CreateTodo) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	Update(id int, todo requests.UpdateTodo) (models.Todo, bool)
	Delete(id int) (bool, interface{})
	GetOne(id int) (models.Todo, error)
}

type todoService struct {
	todoRepo repositories.TodoRepository
}

func NewTodoService(todoRepo repositories.TodoRepository) TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t *todoService) Create(todo requests.CreateTodo) (models.Todo, error) {
	createTodo := models.Todo{}

	err := smapping.FillStruct(&createTodo, smapping.MapFields(&todo))

	if err != nil {
		log.Println("Smapping ", err)
		return models.Todo{}, err
	}

	result, errCreate := t.todoRepo.Create(createTodo)

	if errCreate != nil {
		return models.Todo{}, errCreate
	}

	return result, nil
}

func (t *todoService) GetAll() ([]models.Todo, error) {
	todo, err := t.todoRepo.GetAll()
	if err != nil {
		return nil, err

	}

	return todo, nil

}

func (t *todoService) Update(id int, todo requests.UpdateTodo) (models.Todo, bool) {
	todoUpdate := models.Todo{}

	err := smapping.FillStruct(&todoUpdate, smapping.MapFields(&todo))

	if err != nil {
		log.Println("[ActivityServiceImpl.Create] error fill struct", err)
		return models.Todo{}, false
	}

	update, status := t.todoRepo.Update(id, todoUpdate)

	if !status {
		log.Println("Service Update Activity ", err)
		return models.Todo{}, status
	}

	return update, true
}

func (t *todoService) Delete(id int) (bool, interface{}) {
	delete, err := t.todoRepo.Delete(id)

	if err != nil {
		return false, err
	}

	return delete, nil
}

func (t *todoService) GetOne(id int) (models.Todo, error) {
	todo, err := t.todoRepo.GetOne(id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}
