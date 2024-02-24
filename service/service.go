package service

type Service struct {
	RequestId string
}

func New(requestId string) Service {
	serviceObj := Service{RequestId: requestId}
	return serviceObj
}
