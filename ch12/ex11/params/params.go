package params

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Pack encode value to URL query parameter string
func Pack(v interface{}) string {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return ""
	}

	var buf bytes.Buffer
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		name := field.Tag.Get("http")
		if name == "" {
			name = strings.ToLower(field.Name)
		}

		value := val.Field(i)

		if buf.Len() != 0 {
			buf.WriteRune('&')
		}

		fmt.Fprintf(&buf, "%v=%v", name, value)
	}

	return buf.String()
}
