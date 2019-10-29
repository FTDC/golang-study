package commontest

import (
	"encoding/json"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestStruct2Json(t *testing.T) {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
	var person Person
	json.Unmarshal([]byte(jsonStr), &person)
	t.Log(person)
}
