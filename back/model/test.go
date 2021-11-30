package model

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// checkNotNullFields함수 테스트용 모델
type Test struct {
	Bool       bool
	Int        int
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	Byte       byte
	String     string
	Time       time.Time
	Slice      []string
	Map        map[string]string
	Interface  interface{}
	Dummy      string
}

func (test *Test) NotNullFields() []interface{} {
	return []interface{}{
		test.Bool,
		test.Int,
		test.Int8,
		test.Int16,
		test.Int32,
		test.Int64,
		test.Float32,
		test.Float64,
		test.Complex64,
		test.Complex128,
		test.Byte,
		test.String,
		test.Time,
		test.Slice,
		test.Map,
		test.Interface,
	}
}

func (test *Test) IndexFields() []mongo.IndexModel {
	return nil
}

func (test *Test) SetMaketime() {

}
