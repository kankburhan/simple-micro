package repository

import (
	"errors"
	"simple-micro/app/model/entity"

	"gorm.io/gorm"
)

type todoRepo struct {
	br BaseRepository
}

type TodoRepository interface {
	FindById(id *int) (*entity.Todo, error)
	Find() ([]entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id *int, input map[string]interface{}) error
	Delete(id int) error
}

func NewTodoRepository(base BaseRepository) TodoRepository {
	return &todoRepo{base}
}

func (r *todoRepo) FindById(id *int) (*entity.Todo, error) {
	var todo entity.Todo
	result, err := r.br.FindByid(*id, todo)
	return result.(*entity.Todo), err
}

func (r *todoRepo) Find() ([]entity.Todo, error) {
	var todos []entity.Todo

	err := r.br.GetDB().
		Find(&todos).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return todos, nil
}

func (r *todoRepo) Create(todo *entity.Todo) (*entity.Todo, error) {
	result, err := r.br.Create(&todo)
	return result.(*entity.Todo), err
}

func (r *todoRepo) Update(id *int, input map[string]interface{}) error {
	return r.br.UpdateByid(*id, input, &entity.Todo{})
}

func (r *todoRepo) Delete(id int) error {
	var todo entity.Todo
	return r.br.DeleteByid(id, todo)
}
