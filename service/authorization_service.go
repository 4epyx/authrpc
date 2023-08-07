package service

import (
	"context"
	"errors"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/util"
	"github.com/rs/zerolog"
)

type AuthorizationService struct {
	log zerolog.Logger
	pb.UnimplementedAuthorizationServiceServer
}

func NewAuthorizationService(logger zerolog.Logger) *AuthorizationService {
	return &AuthorizationService{
		log: logger,
	}
}

func (s *AuthorizationService) AuthorizeUser(ctx context.Context, in *pb.Empty) (*pb.AuthUserData, error) {
	token, err := util.GetAuthorizationToken(ctx)
	if err != nil {
		s.log.Info().Str("method", "AuthorizeUser").Err(err).Send()
		return nil, err
	}

	claims, err := util.GetJWTClaims(token, util.JwtSecret)
	if err != nil {
		s.log.Error().Str("method", "AuthorizeUser").Err(err).Send()
		return nil, err
	}

	userIdFloat, ok := claims["user_id"].(float64)
	if !ok {
		s.log.Error().Str("method", "AuthorizeUser").Str("error", "can not parse user id from token").Send()
		return nil, errors.New("can not parse user id from token")
	}
	userId := int64(userIdFloat)

	s.log.Info().Str("method", "AuthorizeUser").Int64("user_id", userId).Msg("authorized")
	return &pb.AuthUserData{
		Id:    userId,
		Email: claims["user_email"].(string),
	}, err
}
