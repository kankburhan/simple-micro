package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-micro/app/api/endpoint"
	"simple-micro/app/model/request"
	"simple-micro/app/service"
	"simple-micro/helper/encoder"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func TodoHttpHandler(s service.TodoService, logger log.Logger) http.Handler {
	pr := mux.NewRouter()

	ep := endpoint.MakeTodoEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encoder.EncodeError),
	}

	pr.Methods("POST").Path("/api/todo").Handler(httptransport.NewServer(
		ep.Add,
		decodeSave,
		encoder.EncodeResponseHTTP,
		options...,
	))

	pr.Methods("GET").Path("/api/todo").Handler(httptransport.NewServer(
		ep.List,
		decodeShowProduct,
		encoder.EncodeResponseHTTP,
		options...,
	))

	pr.Methods("GET").Path("/todo/{id}").Handler(httptransport.NewServer(
		ep.Detail,
		decodeShowProduct,
		encoder.EncodeResponseHTTP,
		options...,
	))

	return pr
}

func decodeSave(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req request.SaveTodo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeShowProduct(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	uid := mux.Vars(r)["id"]
	return uid, nil
}
