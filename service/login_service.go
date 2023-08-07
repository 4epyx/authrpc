package service

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repository"
	"github.com/4epyx/authrpc/util"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	userRepository repository.UserRepository
	pb.UnimplementedLoginServiceServer
}

func NewLoginService(userRepository repository.UserRepository) *LoginService {
	return &LoginService{userRepository: userRepository}
}

func (s *LoginService) LoginUser(ctx context.Context, in *pb.LoginRequest) (*pb.AccessToken, error) {
	user, err := s.userRepository.GetUserLoginData(context.TODO(), in.Login)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, err
	}

	token, err := util.GenerateUserAccessToken(user, util.JwtSecret)
	return &pb.AccessToken{AccessToken: token}, err
}
