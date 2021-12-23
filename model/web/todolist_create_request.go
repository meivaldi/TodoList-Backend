package web

type TodoListCreateRequest struct {
	Title       string `validate:"required,max=255,min=1" json:"title"`
	Description string `validate:"required,max=1000,min=1" json:"desc"`
	Thumbnail   string `validate:"required",max=255,min=10 json:"thumbnail"`
	Priority    int    `json: "priority"`
}
