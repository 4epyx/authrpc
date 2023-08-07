package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repository"
	"github.com/4epyx/authrpc/util"
	"github.com/rs/zerolog"
)

type UserDataService struct {
	userRepository repository.UserRepository
	log            zerolog.Logger
	pb.UnimplementedUserDataServiceServer
}

func NewUserDataService(userRepository repository.UserRepository, logger zerolog.Logger) *UserDataService {
	return &UserDataService{
		userRepository: userRepository,
		log:            logger,
	}
}

func (s *UserDataService) GetCurrentUserData(ctx context.Context, e *pb.Empty) (*pb.User, error) {
	token, err := util.GetAuthorizationToken(ctx)
	if err != nil {
		s.log.Info().Err(err).Send()
		return nil, err
	}

	claims, err := util.GetJWTClaims(token, util.JwtSecret)
	if err != nil {
		s.log.Error().Str("method", "GetCurrentUserData").Err(err).Send()
		return nil, err
	}

	userIdFloat, ok := claims["user_id"].(float64)
	if !ok {
		s.log.Error().Str("method", "GetCurrentUserData").Str("error", "can not parse user id from token").Send()
		return nil, errors.New("can not parse user id from token")
	}
	userId := int64(userIdFloat)

	user, err := s.userRepository.GetUserDataById(ctx, userId)
	if err != nil {
		s.log.Error().Str("method", "GetCurrentUserData").Err(err).Send()
	}

	s.log.Info().Str("method", "GetCurrentUserData").Int64("user_id", user.Id).Msg(fmt.Sprintf("user with id %d gets himself data"))
	return user, nil
}

func (s *UserDataService) GetOtherUserData(ctx context.Context, userId *pb.UserId) (*pb.OtherUser, error) {
	user, err := s.userRepository.GetPublicUserDataById(ctx, userId.Id)
	if err != nil {
		s.log.Info().Str("method", "GetOtherUserData").Err(err).Send()
	}

	return user, nil
}
