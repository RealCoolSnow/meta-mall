package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "user-service/api/user/v1"
	"user-service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	pb.UnimplementedUserServiceServer
	log  *log.Helper
	user *biz.UserUseCase
}
