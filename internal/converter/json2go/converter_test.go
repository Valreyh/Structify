package json2go

import (
	"testing"
)

func TestConvertSimple(t *testing.T) {
	jsonInput := `{"name": "John", "age": 30}`
	expected := "package main\n\ntype Person struct {\n    Age float64 `json:\"age\"`\n    Name string `json:\"name\"`\n}\n"

	result, err := Convert([]byte(jsonInput), "Person")
	if err != nil {
		t.Fatalf("Convert failed: %v", err)
	}

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}
