package services

import (
	"context"
	"os"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/utils"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type LoginService struct {
	db *pgx.ConnPool
	pb.UnimplementedLoginServiceServer
}

func (s *LoginService) LoginUser(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.AccessToken, error) {
	user := pb.User{}
	if err := s.db.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1 OR email = $1",
		in.Login).Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, err
	}

	token, err := utils.GenerateUserAccessToken(&user, os.Getenv("JWT_SECRET"))
	return &pb.AccessToken{AccessToken: token}, err
}
