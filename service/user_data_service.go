package service

import (
	"context"
	"fmt"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repository"
	"github.com/4epyx/authrpc/util"
)

type UserDataService struct {
	userRepository repository.UserRepository
	pb.UnimplementedUserDataServiceServer
}

func NewUserDataService(userRepository repository.UserRepository) *UserDataService {
	return &UserDataService{userRepository: userRepository}
}

func (s *UserDataService) GetCurrentUserData(ctx context.Context, e *pb.Empty) (*pb.User, error) {
	token, err := util.GetAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := util.GetJWTClaims(token, util.JwtSecret)
	if err != nil {
		return nil, err
	}

	userIdFloat, ok := claims["user_id"].(float64)
	if !ok {
		return nil, fmt.Errorf("can not parse user id from token")
	}
	userId := int64(userIdFloat)

	return s.userRepository.GetUserDataById(ctx, userId)
}

func (s *UserDataService) GetOtherUserData(ctx context.Context, userId *pb.UserId) (*pb.OtherUser, error) {
	return s.userRepository.GetPublicUserDataById(ctx, userId.Id)
}
