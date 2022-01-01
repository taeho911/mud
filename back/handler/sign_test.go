package handler

import (
	"context"
	"taeho/mud/agent"
	"taeho/mud/model"
	"taeho/mud/utils"
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
			Password: "noNumber%",
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

func TestIsCorrectPwd(t *testing.T) {
	ctx := context.TODO()
	salt, err := utils.MakeRandom(16)
	if err != nil {
		t.Fatalf("utils.MakeRandom(16) failed. err = %v", err)
	}
	insertPwd := []byte("TestIsCorrectPwd")
	hashedPwd, err := utils.HashPwd(insertPwd, salt)
	if err != nil {
		t.Fatalf("utils.HashPwd(insertPwd, salt) failed. err = %v", err)
	}
	insertUser := model.User{
		Username: "TestIsCorrectPwd",
		Password: utils.EncodeBase64(hashedPwd),
	}
	insertSalt := model.Salt{
		Username: "TestIsCorrectPwd",
		Salt:     salt,
	}
	agent.UserInsertOne(ctx, &insertUser)
	agent.SaltInsertOne(ctx, &insertSalt)

	testcase := []model.User{
		{
			Username: "TestIsCorrectPwd",
			Password: "TestIsCorrectPwd",
		},
		{
			Username: "TestIsCorrectPwd",
			Password: "wrongpassword",
		},
	}
	for i, v := range testcase {
		switch i {
		case 0:
			if err := isCorrectPwd(ctx, v); err != nil {
				t.Fatalf("test case failed. i = %v, err = %v", i, err)
			}
		case 1:
			if err := isCorrectPwd(ctx, v); err == nil {
				t.Fatalf("test case failed. i = %v, err = %v", i, err)
			}
		}
	}

	agent.UserDeleteByID(ctx, insertUser.ID)
	agent.SaltDeleteByID(ctx, insertSalt.ID)
}
