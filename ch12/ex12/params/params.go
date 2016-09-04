package params

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
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

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(form url.Values, ptr interface{}) error {
	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	formats := make(map[string]string)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		var format string
		names := strings.Split(name, ",")
		if len(names) == 2 {
			name = names[0]
			format = names[1]
		}
		fields[name] = v.Field(i)
		formats[name] = format
	}

	// Update struct field for each parameter in the request.
	for name, values := range form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if format, ok := formats[name]; ok && format != "" {
				if validate, ok := validators[format]; ok {
					if !validate(reflect.ValueOf(value)) {
						return fmt.Errorf("failed to parse %v as %s", value, format)
					}
				} else {
					return fmt.Errorf("unknown format: %s", format)
				}
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

// Validate values while unpacking form.
type Validate func(v reflect.Value) (ok bool)

var validators = make(map[string]Validate)

// RegisterValidator register new validation function.
func RegisterValidator(name string, f Validate) {
	validators[name] = f
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
