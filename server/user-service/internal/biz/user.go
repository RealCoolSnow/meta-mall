package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	ID         int64
	GroupId    int32
	UserName   string
	PassWord   string
	Mobile     string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
}
type UserRepo interface {
	// db
	ListUser(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user *User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) List(ctx context.Context) (users []*User, err error) {
	users, err = uc.repo.ListUser(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (user *User, err error) {
	user, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) Update(ctx context.Context, id int64, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *UserUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteUser(ctx, id)
}
