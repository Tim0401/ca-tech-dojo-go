package input

type CreateUser struct {
	Name string `json:"name"`
}

type GetUser struct {
	Xtoken string `json:"x-token"`
}

type UpdateUser struct {
	Xtoken string `json:"x-token"`
	Name   string `json:"name"`
}
