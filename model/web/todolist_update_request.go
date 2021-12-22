package web

type TodoListUpdateRequest struct {
	Id          int
	Title       string
	Description string
	Thumbnail   string
	Priority    int
}
