package domain

type user struct {
	Name         string
	Phone_number string
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       int
	Category    string
	authorID    int
}
