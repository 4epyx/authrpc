package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/utils"
	"github.com/jackc/pgx"
	"google.golang.org/grpc/metadata"
)

type UserDataService struct {
	db *pgx.ConnPool
	pb.UnimplementedUserDataServiceServer
}

func (s *UserDataService) GetCurrentUserData(ctx context.Context, e *pb.Empty) (*pb.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("error while getting metadata from context")
	}

	auth := md["Authorization"]
	if len(auth) == 0 {
		return nil, fmt.Errorf("unauthorized")
	}

	token := strings.TrimPrefix(auth[0], "Token ")
	claims, err := utils.GetJWTClaims(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	userId, ok := claims["user_id"].(int64)
	if !ok {
		return nil, fmt.Errorf("can not parse user id from token")
	}

	user := &pb.User{}
	if err := s.db.QueryRow("SELECT id, email, username FROM users WHERE id = $1", userId).Scan(user.Id, user.Email, user.Username); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserDataService) GetOtherUserData(ctx context.Context, userId *pb.UserId) (*pb.OtherUser, error) {
	user := &pb.OtherUser{}

	if err := s.db.QueryRow("SELECT username FROM user WHERE id = $1", userId.Id).Scan(user.Username); err != nil {
		return nil, err
	}

	user.Id = userId.Id

	return user, nil
}
