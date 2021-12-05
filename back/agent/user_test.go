package agent

import (
	"context"
	"taeho/mud/model"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestUserInsertOne(t *testing.T) {
	ctx := context.TODO()
	entity := model.User{
		Username: "TestUserInsertOne",
		Password: "TestUserInsertOne",
	}
	entity2 := entity
	err := UserInsertOne(ctx, &entity)
	if err != nil {
		t.Fatalf("TestUserInsertOne Fail. err = %v", err)
	}
	t.Log("entity ID =", entity.ID)
	t.Log("entity2 ID =", entity2.ID)
	if entity.Username != entity2.Username {
		t.Fatalf("result.Username != entity.Username. entity.Username = %v", entity.Username)
	}
	if entity.Password != entity2.Password {
		t.Fatalf("result.Password != entity.Password. entity.Password = %v", entity.Password)
	}
	UserDeleteByID(ctx, entity.ID)
}

func TestUserFindByUsername(t *testing.T) {
	ctx := context.TODO()
	result := model.User{
		Username: "TestUserFindByUsername",
		Password: "TestUserFindByUsername",
	}
	UserInsertOne(ctx, &result)

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

	findResult2, err := UserFindByUsername(ctx, "NonExistUsername")
	if err != mongo.ErrNoDocuments {
		t.Logf("findResult2 = %v", findResult2)
		t.Fatalf("TestUserFindByUsername Fail. err = %v", err)
	}
}
