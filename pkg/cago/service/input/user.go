package input

type CreateUser struct {
	Name string
}

type GetUser struct {
	ID int
}

type UpdateUser struct {
	ID   int
	Name string
}
