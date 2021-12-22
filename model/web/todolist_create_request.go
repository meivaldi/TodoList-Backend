package web

type TodoListCreateRequest struct {
	Title       string
	Description string
	Thumbnail   string
	Priority    int
}
