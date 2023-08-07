package repository

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/util"
)

type UserRepository interface {
	CreateUser(context.Context, *pb.RegisterUserRequest) error
	GetUserLoginData(context.Context, string) (*util.User, error)
	GetUserDataById(context.Context, int64) (*pb.User, error)
	GetPublicUserDataById(context.Context, int64) (*pb.OtherUser, error)
}
