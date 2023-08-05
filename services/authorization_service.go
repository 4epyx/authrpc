package services

import (
	"context"
	"os"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/utils"
)

type AuthorizationService struct {
	pb.UnimplementedAuthorizationServiceServer
}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

func (s *AuthorizationService) IsAuthorized(ctx context.Context, in *pb.AccessToken) (*pb.BoolResponse, error) {
	_, err := utils.GetJWTClaims(in.AccessToken, os.Getenv("JWT_SECRET"))

	return &pb.BoolResponse{Flag: err == nil}, err
}
