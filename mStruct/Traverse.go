package mStruct

import (
	"reflect"
)

/*
	RangeStruct(opt, func(key string, val any) {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   key,
					Value: val,
				},
			},
		})
	})
*/

func Traverse(opt any, callback func(key string, val any)) {
	value := reflect.ValueOf(opt)
	lType := reflect.TypeOf(opt)

	for i := 0; i < lType.NumField(); i++ {
		name := lType.Field(i).Name
		val := value.Field(i).Interface()

		callback(name, val)
	}
}
