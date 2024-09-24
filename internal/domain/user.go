package domain

import (
	"clean-arch/types"
	"context"
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID       int  `json:"id"`
	Metadata Meta `json:"meta"`
	IsAdmin  bool `json:"is_admin"`
	Profile
}

func (u *User) Validate() error {
	return v.ValidateStruct(u,
		v.Field(&u.ID, v.Required),
	)
}

type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Meta struct {
	Method      string  `json:"method"`
	URI         string  `json:"uri"`
	ServiceName *string `json:"serviceName,omitempty"`
	AppKey      *string `json:"app-key,omitempty"`
	Profile
	Payload interface{} `json:"payload"`
}

type UserUseCase interface {
	CreateUser(ctx context.Context, req types.UserReq) (*types.UserResp, error)
}

type UserRepo interface {
	CreateUser(req User) (*types.UserResp, error)
}
