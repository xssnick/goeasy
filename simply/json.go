package simply

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/xssnick/goeasy"
	"github.com/xssnick/goeasy/flexy"
)

type Dynamic = map[string]interface{}

func Empty() interface{} {
	return struct{}{}
}

func Json(entity interface{}, data []byte) interface{} {
	p := reflect.New(reflect.TypeOf(entity))

	err := json.Unmarshal(data, p.Interface())
	if err != nil {
		log.Println(err)
		return goeasy.NewBasicError(goeasy.ErrCodeBadRequest, "invalid json")
	}

	return p.Elem().Interface()
}

func JsonPtr(entity interface{}, data []byte) interface{} {
	err := json.Unmarshal(data, entity)
	if err != nil {
		log.Println(err)
		return goeasy.NewBasicError(goeasy.ErrCodeBadRequest, "invalid json")
	}

	return entity
}

func JsonF(entity interface{}, data []byte) interface{} {
	p := flexy.ToPointerValue(entity)

	err := json.Unmarshal(data, flexy.FastPack(p))
	if err != nil {
		log.Println(err)
		return goeasy.NewBasicError(goeasy.ErrCodeBadRequest, "invalid json")
	}

	return flexy.FastPack(p.Elem())
}
