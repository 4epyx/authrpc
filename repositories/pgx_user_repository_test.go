package repositories_test

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/repositories"
	"github.com/4epyx/authrpc/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

var testUser *pb.RegisterUserRequest = &pb.RegisterUserRequest{
	Email:    "aboba@gmail.com",
	Username: "aboba",
	Password: "verystrongpasword228",
}

func setupDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	db, err := utils.ConnectToDB(ctx, os.Getenv("TEST_DB_URL"))
	if err != nil {
		return nil, err
	}
	if err := utils.MigrateTable(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}

func TestPgxUserRepository_CreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *pb.RegisterUserRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := setupDatabase(ctx)
	if err != nil {
		t.Fatal("failed to connect to database")
	}
	defer db.Close()
	defer db.Exec(ctx, "DROP TABLE users")

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
				ctx:  ctx,
				user: testUser,
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

func TestPgxUserRepository_GetUserLoginData(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := setupDatabase(ctx)
	if err != nil {
		t.Fatal("failed to connect to database")
	}
	defer db.Close()
	defer db.Exec(ctx, "DROP TABLE users")

	if err := utils.MigrateTable(ctx, db); err != nil {
		t.Fatal("failed to migrate table")
	}

	r := repositories.NewPgxUserRepository(db)
	r.CreateUser(ctx, testUser)

	wantUser := &utils.User{
		Id:       1,
		Email:    testUser.Email,
		Password: testUser.Password,
	}

	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name    string
		args    args
		want    *utils.User
		wantErr bool
	}{
		{
			name: "smoke test (get by username)",
			args: args{
				ctx:   ctx,
				login: testUser.Username,
			},
			want:    wantUser,
			wantErr: false,
		},
		{
			name: "smoke test (get by email)",
			args: args{
				ctx:   ctx,
				login: testUser.Email,
			},
			want:    wantUser,
			wantErr: false,
		},
		{
			name: "not existing login",
			args: args{
				ctx:   ctx,
				login: "fake-login",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := r.GetUserLoginData(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgxUserRepository.GetUserLoginData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgxUserRepository.GetUserLoginData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgxUserRepository_GetUserDataById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := setupDatabase(ctx)
	if err != nil {
		t.Fatal("failed to connect to database")
	}
	defer db.Close()
	defer db.Exec(ctx, "DROP TABLE users")

	if err := utils.MigrateTable(ctx, db); err != nil {
		t.Fatal("failed to migrate table")
	}

	r := repositories.NewPgxUserRepository(db)
	r.CreateUser(ctx, testUser)

	wantUser := &pb.User{
		Id:       1,
		Email:    testUser.Email,
		Username: testUser.Username,
	}

	type args struct {
		ctx    context.Context
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.User
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			want:    wantUser,
			wantErr: false,
		},
		{
			name: "not existing user",
			args: args{
				ctx:    ctx,
				userId: 100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repositories.NewPgxUserRepository(db)
			got, err := r.GetUserDataById(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgxUserRepository.GetUserDataById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgxUserRepository.GetUserDataById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgxUserRepository_GetPublicUserDataById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := setupDatabase(ctx)
	if err != nil {
		t.Fatal("failed to connect to database")
	}
	defer db.Close()
	defer db.Exec(ctx, "DROP TABLE users")

	if err := utils.MigrateTable(ctx, db); err != nil {
		t.Fatal("failed to migrate table")
	}

	r := repositories.NewPgxUserRepository(db)
	r.CreateUser(ctx, testUser)

	wantUser := &pb.OtherUser{
		Id:       1,
		Username: testUser.Username,
	}

	type args struct {
		ctx    context.Context
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.OtherUser
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				ctx:    ctx,
				userId: 1,
			},
			want:    wantUser,
			wantErr: false,
		},
		{
			name: "not existing user",
			args: args{
				ctx:    ctx,
				userId: 100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetPublicUserDataById(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgxUserRepository.GetPublicUserDataById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgxUserRepository.GetPublicUserDataById() = %v, want %v", got, tt.want)
			}
		})
	}
}
