package middleware

import (
	"net/http"
)

type Middleware interface {
	exec(http.HandlerFunc) http.HandlerFunc
}

type mwStack struct {
	middlewares []Middleware
}

func NewMws(mws ...Middleware) mwStack {
	return mwStack{append([]Middleware(nil), mws...)}
}

func (m mwStack) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i].exec(h)
	}
	return h
}
