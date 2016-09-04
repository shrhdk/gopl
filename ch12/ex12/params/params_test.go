package params

import (
	"net/url"
	"reflect"
	"regexp"
	"testing"
)

func TestPack(t *testing.T) {
	tests := []struct {
		give interface{}
		want string
	}{
		{"", ""},
		{struct {
			Foo string
			Bar string
		}{"hello", "world"}, "foo=hello&bar=world"},
		{struct {
			Foo int `http:"Bar"`
		}{1}, "Bar=1"},
	}

	for _, test := range tests {
		got := Pack(test.give)
		if got != test.want {
			t.Errorf("want %v, got %v", test.want, got)
		}
	}
}

func TestUnpack(t *testing.T) {
	RegisterValidator("zipcode", func(v reflect.Value) bool {
		if v.Kind() != reflect.String {
			return false
		}
		s := v.Interface().(string)
		return regexp.MustCompile(`^\d{3}-\d{4}$`).MatchString(s)
	})

	type Address struct {
		ZIP string `http:"zip,zipcode"`
	}

	tests := []struct {
		give string
		want bool
	}{
		{"zip=", false},
		{"zip=123", false},
		{"zip=123-45678", false},
		{"zip=123-4567", true},
	}

	for _, test := range tests {
		form, err := url.ParseQuery(test.give)
		if err != nil {
			t.Fatalf("ParseQuery(%s) failed: %v", test.give, err)
		}
		var got Address
		err = Unpack(form, &got)
		if test.want && err != nil {
			t.Errorf("validation failed with valid value %v : %v", test.give, err)
		} else if !test.want && err == nil {
			t.Errorf("validation succeeded with invalid value %v : %v", test.give, err)
		}
	}
}
