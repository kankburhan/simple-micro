package service

import (
	"simple-micro/app/model/entity"
	"simple-micro/app/model/request"
	"simple-micro/app/model/response"
	"simple-micro/app/repository"
)

type TodoService interface {
	Create(input request.SaveTodo) interface{}
	List() interface{}
	Get(id int) interface{}
	Update(id *int, input *request.SaveTodo) interface{}
}

type todoServiceImpl struct {
	brRep   repository.BaseRepository
	todoRep repository.TodoRepository
}

func NewtodoServiceImpl(
	br repository.BaseRepository,
	tr repository.TodoRepository,
) TodoService {
	return &todoServiceImpl{br, tr}
}

func (s *todoServiceImpl) Create(input request.SaveTodo) interface{} {
	s.brRep.BeginTx()
	//Set request to entity
	Todo := entity.Todo{
		Todo: input.Todo,
	}

	result, err := s.todoRep.Create(&Todo)
	if err != nil {
		return response.SetResponse(nil)
	}
	s.brRep.CommitTx()
	return response.SetResponse(result)
}

func (s *todoServiceImpl) Get(id int) interface{} {

	result, err := s.todoRep.FindById(&id)
	if err != nil {
		return response.SetResponse(nil)
	}

	if result == nil {
		return response.SetResponse(nil)
	}

	return response.SetResponse(result)
}

func (s *todoServiceImpl) List() interface{} {
	result, err := s.todoRep.Find()
	if err != nil {
		return response.SetResponse(nil)
	}

	if result == nil {
		return response.SetResponse(nil)
	}

	return response.SetResponse(result)
}

func (s *todoServiceImpl) Update(id *int, input *request.SaveTodo) interface{} {
	_, err := s.todoRep.FindById(id)
	if err != nil {
		return response.SetResponse(nil)
	}

	data := map[string]interface{}{
		"todo": input.Todo,
	}

	err = s.todoRep.Update(id, data)
	if err != nil {
		return response.SetResponse(nil)
	}

	return response.SetResponse(data)
}
