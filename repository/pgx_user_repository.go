package repository

import (
	"context"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/util"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgxUserRepository struct {
	db *pgxpool.Pool
}

func NewPgxUserRepository(db *pgxpool.Pool) *PgxUserRepository {
	return &PgxUserRepository{db: db}
}

func (r *PgxUserRepository) CreateUser(ctx context.Context, user *pb.RegisterUserRequest) error {
	if err := util.ValidateUserData(user); err != nil {
		return err
	}

	_, err := r.db.Exec(ctx, "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)", user.Email, user.Username, user.Password)
	return err
}

func (r *PgxUserRepository) GetUserLoginData(ctx context.Context, login string) (*util.User, error) {
	user := &util.User{}
	if err := r.db.QueryRow(ctx, "SELECT id, email, password FROM users WHERE username = $1 OR email = $1", login).Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PgxUserRepository) GetUserDataById(ctx context.Context, userId int64) (*pb.User, error) {
	user := &pb.User{}
	if err := r.db.QueryRow(ctx, "SELECT id, email, username FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Email, &user.Username); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PgxUserRepository) GetPublicUserDataById(ctx context.Context, userId int64) (*pb.OtherUser, error) {
	user := &pb.OtherUser{}

	if err := r.db.QueryRow(ctx, "SELECT username FROM users WHERE id = $1", userId).Scan(&user.Username); err != nil {
		return nil, err
	}

	user.Id = userId
	return user, nil
}
