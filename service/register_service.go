package services

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

type registerService struct {
	db *pgx.ConnPool
	pb.UnimplementedRegisterServiceServer
}

func NewRegisterService(db *pgx.ConnPool) *registerService {
	return &registerService{db: db}
}

func (s *registerService) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.BoolResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		return &pb.BoolResponse{Flag: false}, err
	}

	_, err = s.db.Exec("INSERT INTO users (email. username, password) VALUES $1, $2, $3", in.Email, in.Username, string(password))

	return &pb.BoolResponse{Flag: err != nil}, err
}
