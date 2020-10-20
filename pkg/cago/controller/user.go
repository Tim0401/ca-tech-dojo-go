package controller

import (
	"ca-tech-dojo-go/pkg/cago/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	us service.UserService
}

// NewUserController ユーザーコントローラー作成
func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// var user model.User
	// ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v\n", jsonBody)

	w.WriteHeader(http.StatusOK)

	// uc.us.CreateUser(ctx, &user)
}
