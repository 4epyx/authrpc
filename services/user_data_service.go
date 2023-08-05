package services

import (
	"context"
	"fmt"
	"os"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repositories"
	"github.com/4epyx/authrpc/utils"
)

type UserDataService struct {
	userRepository repositories.UserRepository
	pb.UnimplementedUserDataServiceServer
}

func NewUserDataService(userRepository repositories.UserRepository) *UserDataService {
	return &UserDataService{userRepository: userRepository}
}

func (s *UserDataService) GetCurrentUserData(ctx context.Context, e *pb.Empty) (*pb.User, error) {
	token, err := utils.GetAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := utils.GetJWTClaims(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	userId, ok := claims["user_id"].(int64)
	if !ok {
		return nil, fmt.Errorf("can not parse user id from token")
	}

	return s.userRepository.GetUserDataById(context.TODO(), userId)
}

func (s *UserDataService) GetOtherUserData(ctx context.Context, userId *pb.UserId) (*pb.OtherUser, error) {
	return s.userRepository.GetPublicUserDataById(context.TODO(), userId.Id)
}
