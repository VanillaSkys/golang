package response

type ResponseData struct {
	Status int32       `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status int32       `json:"status"`
	Error  interface{} `json:"error"`
}
