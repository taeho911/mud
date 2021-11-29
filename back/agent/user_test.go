package agent

import (
	"context"
	"taeho/mud/model"
	"testing"
)

func TestUserInsertOne(t *testing.T) {
	ctx := context.TODO()
	entity := model.User{
		Username: "TestUserInsertOne",
		Password: "TestUserInsertOne",
	}
	result, err := UserInsertOne(ctx, &entity)
	if err != nil {
		t.Fatalf("TestUserInsertOne Fail. err = %v", err)
	}
	t.Log("result ID =", result.ID)
	if result.Username != entity.Username {
		t.Fatalf("result.Username != entity.Username. result.Username = %v", result.Username)
	}
	if result.Password != entity.Password {
		t.Fatalf("result.Password != entity.Password. result.Password = %v", result.Password)
	}
	UserDeleteByID(ctx, result.ID)
}

func TestUserFindByUsername(t *testing.T) {
	ctx := context.TODO()
	entity := model.User{
		Username: "TestUserFindByUsername",
		Password: "TestUserFindByUsername",
	}
	result, _ := UserInsertOne(ctx, &entity)
	findResult, err := UserFindByUsername(ctx, result.Username)
	if err != nil {
		t.Fatalf("TestUserFindByUsername Fail. err = %v", err)
	}
	if findResult.ID != result.ID {
		t.Fatalf("findResult.ID != result.ID. findResult.ID = %v, result.ID = %v", findResult.ID, result.ID)
	}
	if findResult.Username != result.Username {
		t.Fatalf("findResult.Username != result.Username. findResult.Username = %v", findResult.Username)
	}
	if findResult.Password != result.Password {
		t.Fatalf("findResult.Password != result.Password. findResult.Password = %v", findResult.Password)
	}
	UserDeleteByID(ctx, result.ID)
}
