package logic

import "context"

type CreateAccountParams struct {
	Username string
	Password string
}

type CreateSessionParams struct {
	Username string
	Password string
}

type User struct {
	ID       uint64
	Username string
}

type Session struct {
	ID       uint64
	Username string
}

type Account interface {
	CreateAccount(context.Context, CreateAccountParams) (User, error)
	CreateSession(context.Context, CreateSessionParams) (Session, error)
}
