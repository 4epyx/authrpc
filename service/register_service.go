package service

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repository"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	userRepository repository.UserRepository
	log            zerolog.Logger
	pb.UnimplementedRegisterServiceServer
}

func NewRegisterService(userRepository repository.UserRepository, logger zerolog.Logger) *RegisterService {
	return &RegisterService{
		userRepository: userRepository,
		log:            logger,
	}
}

func (s *RegisterService) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.BoolResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		s.log.Error().Str("method", "RegisterUser").Err(err).Send()
		return &pb.BoolResponse{Flag: false}, err
	}
	in.Password = string(password)

	err = s.userRepository.CreateUser(ctx, in)
	if err != nil {
		s.log.Error().Err(err).Send()
		return nil, err
	}

	s.log.Info().Str("method", "RegisterUser").Str("user_email", in.Email).Msg("registered")
	return &pb.BoolResponse{Flag: true}, nil
}
