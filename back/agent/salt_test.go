package agent

import (
	"context"
	"taeho/mud/model"
	"testing"
)

func TestSaltInsertOne(t *testing.T) {
	ctx := context.TODO()
	salt := model.Salt{
		Username: "TestSaltInsertOne",
		Salt:     []byte("TestSaltInsertOne"),
	}
	if err := SaltInsertOne(ctx, &salt); err != nil {
		t.Fatalf("TestSaltInsertOne Fail. err = %v", err)
	}
	SaltDeleteByID(ctx, salt.ID)
}
