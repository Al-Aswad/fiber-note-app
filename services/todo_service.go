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
