package jsoon

import "testing"

func TestJUnmarshal(t *testing.T) {
	type A struct {
		A string `json:"a"`
	}

	a, err := Unmarshal[A]([]byte(`{"a":"xxxx"}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a.A)
	//xxxx
}
func TestJUnmarshalP(t *testing.T) {
	type A struct {
		A string `json:"a"`
	}

	a, err := Unmarshal[*A]([]byte(`{"a":"xxxx"}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a.A)
	//xxxx
}

func TestKeys(t *testing.T) {
	t.Log(KeysMaps("1", "2", "a", "6", "cc"))
	// map[1:map[2:map[a:map[6:map[]]]]]
}

func TestGetNum(t *testing.T) {
	v, ok, err := GetMustAny([]byte(`{"1":2, "2":{"a":{"6":{"cc":1}}}, "3":4}`),
		"2", "a", "6", "cc")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ok, v)
}
