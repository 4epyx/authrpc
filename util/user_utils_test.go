package util_test

import (
	"testing"

	"github.com/4epyx/authrpc/pb"
	"github.com/4epyx/authrpc/util"
)

func TestValidateUserData(t *testing.T) {
	type args struct {
		user *pb.RegisterUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "user without password",
			args: args{
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
				user: &pb.RegisterUserRequest{
					Email:    "aboba5@gmail.com",
					Username: "$aboba5$",
					Password: "verystrongpasword228",
				},
			},
			wantErr: true,
		},
		{
			name: "user with too long username",
			args: args{
				user: &pb.RegisterUserRequest{
					Email:    "aboba5@gmail.com",
					Username: "aboba5fdsasdfdsasadf",
					Password: "verystrongpasword228",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := util.ValidateUserData(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
