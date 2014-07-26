package structure

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToMapNonStruct(t *testing.T) {
	foo := []string{"foo"}

	_, err := ToMap(foo)
	if err == nil {
		t.Error("ToMap shouldn't accept non struct types")
	}

}

func TestToMap(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	a, err := ToMap(T)
	if err != nil {
		t.Error(err)
	}

	if typ := reflect.TypeOf(a).Kind(); typ != reflect.Map {
		t.Errorf("ToMap should return a map type, got: %v", typ)
	}

	// we have three fields
	if len(a) != 3 {
		t.Errorf("ToMap should return a map of len 3, got: %d", len(a))
	}

	inMap := func(val interface{}) bool {
		for _, v := range a {
			if reflect.DeepEqual(v, val) {
				return true
			}
		}

		return false
	}

	for _, val := range []interface{}{"a-value", 2, true} {
		if !inMap(val) {
			t.Errorf("ToMap should have the value %v", val)
		}
	}

}

func TestToMap_Tag(t *testing.T) {
	var T = struct {
		A string `structure:"x"`
		B int    `structure:"y"`
		C bool   `structure:"z"`
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	a, err := ToMap(T)
	if err != nil {
		t.Error(err)
	}

	inMap := func(key interface{}) bool {
		for k := range a {
			if reflect.DeepEqual(k, key) {

				return true
			}
		}

		return false
	}

	for _, key := range []string{"x", "y", "z"} {
		if !inMap(key) {
			t.Errorf("ToMap should have the key %v", key)
		}
	}

}

func ExampleMap() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	s := &Server{
		Name:    "Arslan",
		ID:      123456,
		Enabled: true,
	}

	m, err := ToMap(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", m["Name"])
	fmt.Printf("%#v\n", m["ID"])
	fmt.Printf("%#v\n", m["Enabled"])
	// Output:
	// "Arslan"
	// 123456
	// true

}
