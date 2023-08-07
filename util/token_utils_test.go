package util_test

import (
	"testing"

	"github.com/4epyx/authrpc/util"
)

func TestTokenUtils(t *testing.T) {
	type args struct {
		user      *util.User
		secretKey []byte
	}

	secret := []byte("test_secre.t")

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "smoke test",
			args: args{
				user: &util.User{
					Id:       1,
					Email:    "user@example.com",
					Password: "12345",
				},
				secretKey: secret,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := util.GenerateUserAccessToken(tt.args.user, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateUserAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			decoded, err := util.GetJWTClaims(token, secret)
			if err != nil {
				t.Errorf("GetJWTClaims() error: %v", err)
			}

			if id := int64(decoded["user_id"].(float64)); id != tt.args.user.Id {
				t.Errorf("ids mismatching: user.Id = %d, decoded id = %d", tt.args.user.Id, id)
			}
			if email := decoded["user_email"].(string); email != tt.args.user.Email {
				t.Errorf("emails mismatching: user.Email = %s, decoded email = %s", tt.args.user.Email, email)
			}
		})
	}
}
