package web

type TodoListResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Thumbnail   string `json:"thumbnail"`
	Priority    int    `json:priority`
}
