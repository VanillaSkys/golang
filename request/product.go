package request

type FindByName struct {
	Name string `json:"name"`
}

type CreateProduct struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateById struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DeleteById struct {
	Id string `json:"id"`
}
