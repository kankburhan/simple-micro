package response

type responseHttp struct {
	Data interface{} `json:"data"`
}

func SetResponse(data interface{}) interface{} {
	return responseHttp{
		Data: data,
	}
}
