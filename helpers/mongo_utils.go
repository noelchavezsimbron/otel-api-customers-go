package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func Decimal128ToFloat64(decimal128 primitive.Decimal128) float64 {

	value, err := strconv.ParseFloat(decimal128.String(), 64)
	if err != nil {
		panic(err)
	}
	return value
}

type Aggregate []bson.M
