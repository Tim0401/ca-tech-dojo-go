package input

type CreateUser struct {
	Name string
}

type GetUser struct {
	ID int32
}

type UpdateUser struct {
	ID   int32
	Name string
}
