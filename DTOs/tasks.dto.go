package dtos

type TasksDtoResponse struct {
	ID          uint8  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

type TasksDtoRequest struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=10"`
}
