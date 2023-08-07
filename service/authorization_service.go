package service

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/util"
)

type AuthorizationService struct {
	pb.UnimplementedAuthorizationServiceServer
}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

func (s *AuthorizationService) IsAuthorized(ctx context.Context, in *pb.AccessToken) (*pb.BoolResponse, error) {
	_, err := util.GetJWTClaims(in.AccessToken, util.JwtSecret)

	return &pb.BoolResponse{Flag: err == nil}, err
}
