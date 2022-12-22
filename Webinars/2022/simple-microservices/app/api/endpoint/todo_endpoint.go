package endpoint

import (
	"context"
	"fmt"
	"simple-micro/app/model/request"
	"simple-micro/app/service"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

type TodoEndpoint struct {
	Add    endpoint.Endpoint
	Detail endpoint.Endpoint
	List   endpoint.Endpoint
	Update endpoint.Endpoint
}

func MakeTodoEndpoints(s service.TodoService) TodoEndpoint {
	return TodoEndpoint{
		Add:    makeAdd(s),
		Detail: makeDetail(s),
		List:   makeList(s),
	}
}

func makeAdd(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (response interface{}, err error) {
		req := rqst.(request.SaveTodo)
		return s.Create(req), nil
	}
}

func makeDetail(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (response interface{}, err error) {
		id, err := strconv.Atoi(fmt.Sprint(rqst))
		return s.Get(id), nil
	}
}

func makeList(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (response interface{}, err error) {
		return s.List(), nil
	}
}
