package request

type CreateProduct struct {
	Name string `json:"name" binding:"reuired"`
}
