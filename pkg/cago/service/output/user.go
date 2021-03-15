package output

type CreateUser struct {
	Xtoken string
}
type GetUser struct {
	ID   int
	Name string
}

type UpdateUser struct{}
