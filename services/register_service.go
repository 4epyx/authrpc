package services

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repositories"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	userRepository repositories.UserRepository
	pb.UnimplementedRegisterServiceServer
}

func NewRegisterService(userRepository repositories.UserRepository) *RegisterService {
	return &RegisterService{userRepository: userRepository}
}

func (s *RegisterService) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.BoolResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		return &pb.BoolResponse{Flag: false}, err
	}
	in.Password = string(password)

	err = s.userRepository.CreateUser(ctx, in)

	return &pb.BoolResponse{Flag: err == nil}, err
}
