package web

type TodoListUpdateRequest struct {
	Id          int    `validate="required" json:"id"`
	Title       string `validate="required,min=1,max=255" json:"title"`
	Description string `validate="required,min=1,max=1000 json:"desc"`
	Thumbnail   string `validate="required,min=1,max=255 json:"thumbnail"`
	Priority    int    `json:"priority"`
}
