package json2go

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Convert(jsonData []byte, structName string) (string, error) {
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", fmt.Errorf("invalid JSON: %v", err)
	}

	goCode := generateStruct(structName, data)

	output := fmt.Sprintf("package main\n\n%s", goCode)
	return output, nil
}

func generateStruct(name string, data interface{}) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", name))

	switch v := data.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := v[key]
			fieldName := strings.Title(key)
			fieldType := getGoType(value)
			builder.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))
		}
	default:
		builder.WriteString(fmt.Sprintf("    Value %s `json:\"value\"`\n", getGoType(v)))
	}

	builder.WriteString("}\n")
	return builder.String()
}

// détermine le type Go correspondant à une valeur JSON
func getGoType(value interface{}) string {
	switch v := value.(type) {
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case nil:
		return "interface{}"
	case []interface{}:
		if len(v) > 0 {
			elemType := getGoType(v[0])
			return "[]" + elemType
		}
		return "[]interface{}"
	case map[string]interface{}:
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}
