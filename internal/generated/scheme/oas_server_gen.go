// Code generated by ogen, DO NOT EDIT.

package scheme

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// LoginPost implements POST /login operation.
	//
	// Упрощенный процесс аутентификации путем передачи
	// идентификатор пользователя и получения токена для
	// дальнейшего прохождения авторизации.
	//
	// POST /login
	LoginPost(ctx context.Context, req OptLoginPostReq) (LoginPostRes, error)
	// UserGetIDGet implements GET /user/get/{id} operation.
	//
	// Получение анкеты пользователя.
	//
	// GET /user/get/{id}
	UserGetIDGet(ctx context.Context, params UserGetIDGetParams) (UserGetIDGetRes, error)
	// UserRegisterPost implements POST /user/register operation.
	//
	// Регистрация нового пользователя.
	//
	// POST /user/register
	UserRegisterPost(ctx context.Context, req OptUserRegisterPostReq) (UserRegisterPostRes, error)
	// UserSearchGet implements GET /user/search operation.
	//
	// Поиск анкет.
	//
	// GET /user/search
	UserSearchGet(ctx context.Context, params UserSearchGetParams) (UserSearchGetRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
