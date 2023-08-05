package repositories_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repositories"
	"github.com/4epyx/authrpc/utils"
)

func TestPgxUserRepository_CreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *pb.RegisterUserRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := utils.ConnectToDB(ctx, os.Getenv("TEST_DB_URL"))
	if err != nil {
		t.Fatal("failed to connect to database")
	}
	defer db.Exec(ctx, "DROP TABLE users")
	defer db.Close()

	if err := utils.MigrateTable(ctx, db); err != nil {
		t.Fatal("failed to migrate table")
	}

	r := repositories.NewPgxUserRepository(db)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba@gmail.com",
					Username: "aboba",
					Password: "verystrongpasword228",
				},
			},
			wantErr: false,
		},
		{
			name: "user without password",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba2@gmail.com",
					Username: "aboba2",
				},
			},
			wantErr: true,
		},
		{
			name: "user with incorrect email",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba32examplecom",
					Username: "aboba3",
					Password: "verystrongpasword228",
				},
			},
			wantErr: true,
		},
		{
			name: "user with too short password",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba4@gmail.com",
					Username: "aboba4",
					Password: "ilm",
				},
			},
			wantErr: true,
		},
		{
			name: "user with incorrect username",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba5@gmail.com",
					Username: "$aboba5$",
					Password: "verystrongpasword228",
				},
			},
			wantErr: true,
		},
		{
			name: "user with not unique username",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba6@gmail.com",
					Username: "aboba",
					Password: "verystrongpasword228",
				},
			},
			wantErr: true,
		},
		{
			name: "user with not unique email",
			args: args{
				ctx: ctx,
				user: &pb.RegisterUserRequest{
					Email:    "aboba@gmail.com",
					Username: "aboba7",
					Password: "verystrongpasword228",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("PgxUserRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
