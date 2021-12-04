package handler

import (
	"taeho/mud/model"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	users := []model.User{
		{
			Username: "normal",
		},
		{
			Username: "",
		},
		{
			Username: "white space",
		},
		{
			Username: "usernameovermaximumlength",
		},
	}

	for i, user := range users {
		switch i {
		case 0:
			if err := validateUsername(user); err != nil {
				t.Fatalf("validateUsername failed. err = %v", err)
			}
		default:
			if err := validateUsername(user); err == nil {
				t.Fatalf("validateUsername failed. user.Username = %v", user.Username)
			}
		}
	}
}

func TestValidatePassword(t *testing.T) {
	users := []model.User{
		{
			Password: "norMal!123",
		},
		{
			Password: "",
		},
		{
			Password: "white space",
		},
		{
			Password: "passwordovermaximumlength",
		},
		{
			Password: "nocaptal1@",
		},
		{
			Password: "noNumber%",
		},
		{
			Password: "noSimbol111",
		},
	}

	for i, user := range users {
		switch i {
		case 0:
			if err := validatePassword(user); err != nil {
				t.Fatalf("validatePassword failed. err = %v", err)
			}
		default:
			if err := validatePassword(user); err == nil {
				t.Fatalf("validatePassword failed. user.Password = %v", user.Password)
			}
		}
	}
}
