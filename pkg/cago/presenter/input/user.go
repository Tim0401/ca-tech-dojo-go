package input

type CreateUser struct {
	Xtoken string `json:"token"`
}

type GetUser struct {
	Name string `json:"name"`
}

type UpdateUser struct{}
