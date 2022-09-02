// Package sql  通过一种 转为 string 的数据切片对象,适合多种数据库
package sql

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// 函数更改传入的 string 的值,需要 传入 *string 使用 *s =.... 进行更改
/*
func CCC_C(s *string) {
	b := "b"
	*s = b
}
*/

// NewStringArray 创建一个数据切片字符串对象
func NewStringArray(sArray []string) StringArray {
	return StringArray(strings.Join(sArray, ","))
}

type StringArray string

// Scan implements the sql.Scanner interface.
func (a *StringArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.str(string(src))
	case string:
		return a.str(src)
	case nil:
		return nil
	}

	return fmt.Errorf("cannot convert %T to StringArray", src)
}

func (a *StringArray) Value() (driver.Value, error) {
	return *a, nil
}

func (a *StringArray) str(s string) error {
	*a = StringArray(s)
	return nil
}

func (a *StringArray) String() string {
	if a == nil {
		return ""
	}
	return string(*a)
}

func (a *StringArray) Array() []string {
	if *a == "" {
		return nil
	}
	return strings.Split(a.String(), ",")
}

/*
    common_test.go:31: aaaa
    common_test.go:33: aaaa
    common_test.go:35: aaaa
    common_test.go:38: b



func TestC(t *testing.T) {
	c := "aaaa"
	t.Log(c)
	CCC(c)
	t.Log(c)
	CCC_(&c)
	t.Log(c)
	CCC_C(&c)
	t.Log(c) // b
}
func CCC_C(s *string) {
	b := "b"
	*s = b
}
func CCC_(s *string) {
	b := "b"
	s = &b
}

func CCC(s string) {
	b := "b"
	s = b
}


*/

// Array represents a one-dimensional array of the PostgreSQL/Mysql/Sqlite character types (db type json).
type Array []string

var DefaultNullArray = "[]"

// Scan implements the sql.Scanner interface.
func (a *Array) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scan(src)
	case string:
		return a.scan([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("cannot convert %T to Array", src)
}

func (a *Array) scan(src []byte) error {
	if len(src) <= len([]byte(DefaultNullArray)) {
		*a = nil
		return nil
	}

	b := Array{}
	err := json.Unmarshal(src, &b)
	if err != nil {
		return err
	}
	*a = b
	return nil
}

// Value implements the driver.Valuer interface.
func (a Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	if len(a) > 0 {
		marshal, err := json.Marshal(a)
		if err != nil {
			return nil, err
		}
		return string(marshal), err
	}

	return DefaultNullArray, nil
}
