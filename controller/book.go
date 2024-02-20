package controller

type CreateBookInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateBookOutput struct {
	Name string `json:"name"`
}

func (controller Controller) GetBook() string {
	return "Hello"
}

func (controller Controller) CreateBook(input CreateBookInput) CreateBookOutput {
	output := CreateBookOutput{}
	output.Name = input.Name
	return output
}
