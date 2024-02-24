package service

type Service struct {
	RequestId string
}

func New(requestId string) Service {
	return Service{RequestId: requestId}
}
