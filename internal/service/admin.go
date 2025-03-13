package service

import (
	"context"
	"errors"
	"os"
	"regexp"
	"team_collabrative_dashboard/internal/graph/model"
	dbmodel "team_collabrative_dashboard/internal/model/db_Model"

	"github.com/sirupsen/logrus"
)

func (s *Service) AdminSignUpService(ctx context.Context, input model.AdminInput) (string, error) {
	securityKey := os.Getenv("SECURITY_KEY")
	if securityKey != input.SecurityKey {
		logrus.Error("the security key mismatches")
		return "", errors.New("the secuity key does not match")
	}
	if input.Email == "" {
		logrus.Error("email cannot be empty")
		return "", errors.New("email id cannot be empty")
	} else {
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !re.MatchString(input.Email) {
			logrus.Error("email is not valid")
			return "", errors.New("email id is not a valid emailid")
		}
	}
	if input.Password != input.ConfirmPassword {
		logrus.Error("passwords don't match")
		return "", errors.New("passwords does not match")
	}
	dbData := dbmodel.Admin{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	id, err := s.repo.AdminSignUp(ctx, dbData)
	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}

func AdminSignIN() {

}
