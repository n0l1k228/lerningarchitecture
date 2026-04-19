package dto

type userBase struct {
	Name         string `json:"name" validate:"min=3,max=100"`
	Phone_number string `json:"phone_number" validate:"omitempty,min=8,max=15,startswith=+"`
}

type CreateUserRequest struct {
	userBase
}

type UserResponse struct {
	ID int `json:"id"`
	userBase
}

type ProductDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	AuthorID    int    `json:"author_id"`
}

type ProductSearch struct {
	Title string `json:"title"`
}
