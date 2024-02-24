package repository

type Repository struct {
	RequestId string
}

func New(requestId string) Repository {
	repositoryObj := Repository{RequestId: requestId}
	return repositoryObj
}
