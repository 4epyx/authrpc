package utils

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/4epyx/authrpc/pb"
)

func ValidateUserData(user *pb.RegisterUserRequest) error {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return err
	}

	if strings.ContainsAny(user.Username, " ,.&><\\/!@#$%^*()_-+=[]{}\"'~`") {
		return fmt.Errorf("username cannot contain special characters")
	}
	if len(user.Password) < 4 {
		return fmt.Errorf("password must be at least 4 characters")
	}

	return nil
}
