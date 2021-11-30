package agent

import (
	"context"
	"fmt"
	"os"
	"taeho/mud/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	dbname   string        = "mud"
	minpool  uint64        = 3
	maxpool  uint64        = 7
	connidle time.Duration = 10
	timeout  time.Duration = 3
)

// 실행환경에 맞춰 DB URI를 생성하기 위한 함수
func makeDatabaseURI() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "27017"
	}
	if username != "" && password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", username, password, host, port)
	} else {
		return fmt.Sprintf("mongodb://%s:%s/", host, port)
	}
}

// 메인 모듈에서 http 리스닝 이전에 DB client를 만들어 놓기 위한 함수
func CreateClient(ctx context.Context) error {
	options := options.Client().ApplyURI(makeDatabaseURI())
	options.SetMinPoolSize(minpool)
	options.SetMaxPoolSize(maxpool)
	options.SetMaxConnIdleTime(connidle)
	// mongodb의 커넥션이 증가 혹은 감소할 때 실행중인 세션의 수를 출력한다.
	// options.SetPoolMonitor(&event.PoolMonitor{
	// 	Event: func(evt *event.PoolEvent) {
	// 		switch evt.Type {
	// 		case event.GetSucceeded:
	// 			log.Println("DB Conn++ :", client.NumberSessionsInProgress())
	// 		case event.ConnectionReturned:
	// 			log.Println("DB Conn-- :", client.NumberSessionsInProgress())
	// 		}
	// 	},
	// })
	var err error
	client, err = mongo.NewClient(options)
	if err != nil {
		return err
	}
	if err := client.Connect(ctx); err != nil {
		return err
	}
	return nil
}

func DeleteClient(ctx context.Context) {
	client.Disconnect(ctx)
}

func getColl(collname string) *mongo.Collection {
	return client.Database(dbname).Collection(collname)
}

// 각 콜렉션별로 필요한 인덱스들을 모델에서 추출하여 작성한다.
func createIndexes(collname string, indexModels []mongo.IndexModel) ([]string, error) {
	name, err := getColl(collname).Indexes().CreateMany(context.TODO(), indexModels, nil)
	if err != nil {
		return nil, err
	}
	return name, nil
}

// 모델에서 작성된 NotNull 필드들을 검사하여 디폴트 값이 들어가 있을 경우 error를 반환한다.
func checkNotNullFields(entity model.Model) error {
	notNullFields := entity.NotNullFields()
	for _, value := range notNullFields {
		switch v := value.(type) {
		case bool:
			if !v {
				return fmt.Errorf("NotNullFields cannot have default value. type=%t, value=%v", v, v)
			}
		case int, int8, int16, int32, int64:
			if v == 0 {
				return fmt.Errorf("NotNullFields cannot have default value. type=%t, value=%v", v, v)
			}
		case float32, float64:
			if v == 0 {
				return fmt.Errorf("NotNullFields cannot have default value. type=%t, value=%v", v, v)
			}
		case complex64, complex128:
			if v == 0 {
				return fmt.Errorf("NotNullFields cannot have default value. type=%t, value=%v", v, v)
			}
		case byte:
			if v == byte(0) {
				return fmt.Errorf("NotNullFields cannot have default value. type=byte, value=%v", v)
			}
		case string:
			if len(v) == 0 {
				return fmt.Errorf("NotNullFields cannot have default value. type=string, value=%v", v)
			}
		case time.Time:
			if v.IsZero() {
				return fmt.Errorf("NotNullFields cannot have default value. type=time.Time, value=%v", v)
			}
		default:
			if v == nil {
				return fmt.Errorf("NotNullFields cannot have default value. type=%v, value=%v", v, v)
			}
		}
	}
	return nil
}

func insertOne(collname string, entity model.Model, ctx context.Context, option *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	entity.SetMaketime()
	result, err := getColl(collname).InsertOne(dbctx, entity, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func insertMany(collname string, entity []model.Model, ctx context.Context, option *options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := getColl(collname).InsertMany(dbctx, model.ConvertModelToInterface(entity), option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findOne(collname string, entity interface{}, ctx context.Context, filter interface{}, option *options.FindOneOptions) error {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	var err error
	if option == nil {
		err = getColl(collname).FindOne(dbctx, filter).Decode(entity)
	} else {
		err = getColl(collname).FindOne(dbctx, filter, option).Decode(entity)
	}
	if err != nil {
		return err
	}
	return nil
}

func find(collname string, entity interface{}, ctx context.Context, filter interface{}, option *options.FindOptions) error {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	curosr, err := getColl(collname).Find(dbctx, filter, option)
	if err != nil {
		return err
	}
	defer curosr.Close(dbctx)
	if err := curosr.All(dbctx, entity); err != nil {
		return err
	}
	return nil
}

// update 함수의 주의점
// struct의 필드태그에 omitempty가 들어있지 않을 경우 struct의 초기값이 DB 레코드 값를 덮어쓴다.
// 예시로 foo, bar, baz란 필드를 가지는 struct을 foo만 값을 부여하여 생성한 뒤 update함수에 쏠 경우,
// bar과 baz에는 각 타입별 디폴트 값(int: 0, string: "")이 들어가게 되고, DB 레코드 값을 덮어쓴다.
func updateByID(collname string, id primitive.ObjectID, update model.Model, ctx context.Context, option *options.UpdateOptions) (*mongo.UpdateResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).UpdateByID(dbctx, id, bson.M{"$set": update}, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateOne(collname string, update model.Model, ctx context.Context, filter interface{}, option *options.UpdateOptions) (*mongo.UpdateResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).UpdateOne(dbctx, filter, bson.M{"$set": update}, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func deleteOne(collname string, ctx context.Context, filter interface{}, option *options.DeleteOptions) (*mongo.DeleteResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).DeleteOne(dbctx, filter, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func deleteMany(collname string, ctx context.Context, filter interface{}, option *options.DeleteOptions) (*mongo.DeleteResult, error) {
	dbctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	result, err := getColl(collname).DeleteMany(dbctx, filter, option)
	if err != nil {
		return nil, err
	}
	return result, nil
}
