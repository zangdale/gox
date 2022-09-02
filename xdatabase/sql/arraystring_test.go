package sql

import (
	"reflect"
	"testing"
)

func TestStringArray_Scan(t *testing.T) {
	a := StringArray("string")
	t.Log(a)

	value, err := a.Value()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value, a)

	err = a.Scan("123")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value, a)

	value, err = a.Value()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value, a)

}
func TestArray_Scan(t *testing.T) {
	result, err := Array(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = Array([]string{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `[]`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = Array([]string{`a`, `\b`, `c"`, `d,e`}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `["a","\\b","c\"","d,e"]`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}

}
