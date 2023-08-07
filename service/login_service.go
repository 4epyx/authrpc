package service

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repository"
	"github.com/4epyx/authrpc/util"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	userRepository repository.UserRepository
	log            zerolog.Logger
	pb.UnimplementedLoginServiceServer
}

func NewLoginService(userRepository repository.UserRepository, logger zerolog.Logger) *LoginService {
	return &LoginService{
		userRepository: userRepository,
		log:            logger,
	}
}

func (s *LoginService) LoginUser(ctx context.Context, in *pb.LoginRequest) (*pb.AccessToken, error) {
	user, err := s.userRepository.GetUserLoginData(ctx, in.Login)
	if err != nil {
		s.log.Info().Str("method", "LoginUser").Err(err).Send()
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		s.log.Info().Str("method", "LoginUser").Err(err).Send()
		return nil, err
	}

	token, err := util.GenerateUserAccessToken(user, util.JwtSecret)
	if err != nil {
		s.log.Error().Str("method", "LoginUser").Err(err).Send()
	}
	s.log.Info().Str("method", "LoginUser").Int64("user_id", user.Id).Msg("logged in")

	return &pb.AccessToken{AccessToken: token}, nil
}
