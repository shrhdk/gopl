package params

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Pack encode value to URL query parameter string
func Pack(ptr interface{}) string {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Struct {
		return ""
	}

	var buf bytes.Buffer
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		name := field.Tag.Get("http")
		if name == "" {
			name = strings.ToLower(field.Name)
		}

		value := v.Field(i)

		if buf.Len() != 0 {
			buf.WriteRune('&')
		}

		fmt.Fprintf(&buf, "%v=%v", name, value)
	}

	return buf.String()
}
