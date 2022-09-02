package jsoon

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Unmarshal[T any](data []byte) (T, error) {
	res := new(T)
	err := json.Unmarshal(data, res)
	return *res, err
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32
}

type Float interface {
	float32 | float64
}

type IntUintFloat interface {
	Int | Uint | Float
}

var (
	ErrEmptyKeys = errors.New("empty key string")
)

type MapAny map[string]any

type MapRawMessage map[string]json.RawMessage

func KeysMaps(keys ...string) MapAny {
	tempMap := make(MapAny)
	for i := len(keys) - 2; i > -1; i-- {
		temp := make(MapAny)
		temp[keys[i]] = tempMap
		tempMap = temp
	}
	return tempMap
}

func GetAny(data []byte, keys ...string) (v any, have bool, err error) {
	n := len(keys)
	if n == 0 {
		return 0, false, ErrEmptyKeys
	}

	tempMap := KeysMaps(keys...)
	err = json.Unmarshal(data, &tempMap)
	if err != nil {
		return v, false, err
	}
	for i := 0; i < n-1; i++ {
		if _, ok := tempMap[keys[i]]; ok {
			tempMap, ok = tempMap[keys[i]].(map[string]interface{})
			if !ok {
				return v, false, fmt.Errorf("get keys[%s] is not type map[string]interface{}", keys[:i])
			}
		} else {
			return v, false, fmt.Errorf("get keys[%s] is nil", keys[:i])
		}

	}
	v, have = tempMap[keys[n-1]]
	return v, have, nil
}

func GetMustAny(data []byte, keys ...string) (v any, have bool, err error) {
	n := len(keys)
	if n == 0 {
		return 0, false, ErrEmptyKeys
	}

	tempMap := KeysMaps(keys...)
	err = json.Unmarshal(data, &tempMap)
	if err != nil {
		return v, false, err
	}
	for i := 0; i < n-1; i++ {
		if _, ok := tempMap[keys[i]]; ok {
			tempMap, ok = tempMap[keys[i]].(map[string]interface{})
			if !ok {
				return v, false, nil
			}
		} else {
			return v, false, nil
		}

	}
	v, have = tempMap[keys[n-1]]
	return v, have, nil
}
