package services

import (
	"context"
	"os"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repositories"
	"github.com/4epyx/authrpc/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	userRepository repositories.UserRepository
	pb.UnimplementedLoginServiceServer
}

func NewLoginService(userRepository repositories.UserRepository) *LoginService {
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

	token, err := utils.GenerateUserAccessToken(user, os.Getenv("JWT_SECRET"))
	return &pb.AccessToken{AccessToken: token}, err
}
