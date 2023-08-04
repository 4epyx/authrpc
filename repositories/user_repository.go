package repositories

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/utils"
)

type UserRepository interface {
	CreateUser(context.Context, *pb.RegisterUserRequest) error
	GetUserLoginData(context.Context, string) (*utils.User, error)
	GetUserDataById(context.Context, int64) (*pb.User, error)
	GetPublicUserDataById(context.Context, int64) (*pb.OtherUser, error)
}
