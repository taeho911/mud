package agent

import (
	"context"
	"taeho/mud/model"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMoneyFindByTagsIn(t *testing.T) {
	ctx := context.TODO()
	username := "TestMoneyFindByTagsIn"
	moneyList := []model.Money{
		{
			Username: username,
			Date:     time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			Amount:   1000,
			Tags:     []string{"foo", "bar", "baz"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 2, 12, 0, 0, 0, 0, time.UTC),
			Amount:   -1000,
			Tags:     []string{"bar", "baz"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 3, 12, 0, 0, 0, 0, time.UTC),
			Amount:   1000000000,
			Tags:     []string{"foo", "baz"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 4, 12, 0, 0, 0, 0, time.UTC),
			Amount:   10,
			Tags:     []string{"foo"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 5, 12, 0, 0, 0, 0, time.UTC),
			Amount:   -1000000000,
			Tags:     []string{"foo", "bar"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 6, 12, 0, 0, 0, 0, time.UTC),
			Amount:   10000,
			Tags:     []string{"bar"},
		},
	}

	var deleteList []primitive.ObjectID

	for _, money := range moneyList {
		MoneyInsertOne(ctx, &money)
		deleteList = append(deleteList, money.ID)
	}

	result, err := MoneyFindByTagsIn(ctx, username, []string{"foo"})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(result) != 4 {
		t.Fatalf("len(result) = %v", len(result))
	}

	result2, err := MoneyFindByTagsIn(ctx, username, []string{"foo", "bar"})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(result2) != 6 {
		t.Fatalf("len(result2) = %v", len(result2))
	}

	result3, err := MoneyFindByTagsIn(ctx, username, []string{"dummy"})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(result3) != 0 {
		t.Fatalf("len(result3) = %v", len(result3))
	}

	for _, id := range deleteList {
		MoneyDeleteByID(ctx, id, username)
	}
}

func TestMoneyFindByTagsAll(t *testing.T) {
	ctx := context.TODO()
	username := "TestMoneyFindByTagsAll"
	moneyList := []model.Money{
		{
			Username: username,
			Date:     time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			Amount:   1000,
			Tags:     []string{"foo", "bar", "baz"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 2, 12, 0, 0, 0, 0, time.UTC),
			Amount:   -1000,
			Tags:     []string{"bar", "baz"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 3, 12, 0, 0, 0, 0, time.UTC),
			Amount:   1000000000,
			Tags:     []string{"foo", "baz"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 4, 12, 0, 0, 0, 0, time.UTC),
			Amount:   10,
			Tags:     []string{"foo"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 5, 12, 0, 0, 0, 0, time.UTC),
			Amount:   -1000000000,
			Tags:     []string{"foo", "bar"},
		},
		{
			Username: username,
			Date:     time.Date(2021, 6, 12, 0, 0, 0, 0, time.UTC),
			Amount:   10000,
			Tags:     []string{"bar"},
		},
	}

	var deleteList []primitive.ObjectID
	MoneyDeleteByUsername(ctx, username)

	for _, money := range moneyList {
		MoneyInsertOne(ctx, &money)
		deleteList = append(deleteList, money.ID)
	}

	result, err := MoneyFindByTagsAll(ctx, username, []string{"foo", "bar", "baz"})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("len(result) = %v", len(result))
	}

	for _, id := range deleteList {
		MoneyDeleteByID(ctx, id, username)
	}
}

func TestMoneyFindByMonth(t *testing.T) {
	ctx := context.TODO()
	username := "TestMoneyFindByMonth"
	moneyList := []model.Money{
		{
			Username: username,
			Date:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			Amount:   1000,
		},
		{
			Username: username,
			Date:     time.Date(2021, 1, 31, 23, 59, 59, 0, time.UTC),
			Amount:   -1000,
		},
		{
			Username: username,
			Date:     time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			Amount:   1000,
		},
		{
			Username: username,
			Date:     time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			Amount:   1000,
		},
		{
			Username: username,
			Date:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			Amount:   1000,
		},
	}

	var deleteList []primitive.ObjectID
	MoneyDeleteByUsername(ctx, username)

	for _, money := range moneyList {
		MoneyInsertOne(ctx, &money)
		deleteList = append(deleteList, money.ID)
	}

	result1, err := MoneyFindByMonth(ctx, username, 2021, 1, 1)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(result1) != 2 {
		t.Fatalf("len(result1) = %v", len(result1))
	}

	result2, err := MoneyFindByMonth(ctx, username, 2021, 12, 1)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if len(result2) != 1 {
		t.Fatalf("len(result2) = %v", len(result2))
	}

	for _, id := range deleteList {
		MoneyDeleteByID(ctx, id, username)
	}
}
