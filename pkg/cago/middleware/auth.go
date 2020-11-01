package middleware

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/repository"
	"context"
	"fmt"
	"net/http"
)

type authMiddleware struct {
	r repository.Repository
}

// NewAuthMiddleware ユーザーRouter作成
func NewAuthMiddleware(r repository.Repository) Middleware {
	return &authMiddleware{r}
}

func (am *authMiddleware) exec(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] authMiddleware")

		// ユーザー認証
		con, err := am.r.NewConnection()
		if err != nil {
			panic(err)
		}
		defer con.Close()

		token := r.Header.Get("x-token")
		modelUser, err := con.User().FindByToken(token)
		if err != nil {
			http.Error(w, "auth failed", http.StatusUnauthorized)
			return
		}

		// 格納
		ctx := r.Context()
		sctx := context.WithValue(ctx, model.UserKey, modelUser)
		r = r.WithContext(sctx)
		next.ServeHTTP(w, r)

		fmt.Println("[END] authMiddleware")
	}
}
