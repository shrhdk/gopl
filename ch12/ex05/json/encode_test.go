package json

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		given interface{}
		want  string
	}{
		{nil, "null"},
		{true, "true"},
		{false, "false"},
		{map[string]string{"hello": "world", "foo": "bar"}, `{"hello":"world","foo":"bar"}`},
		{map[string]int{"foo": 0, "bar": 1}, `{"foo":0,"bar":1}`},
		{struct {
			Hello int
			World int
		}{0, 1}, `{"Hello":0,"World":1}`},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		err := encode(&buf, reflect.ValueOf(test.given))
		if err != nil {
			t.Errorf("error occured while encoding %v. %v", test.given, err)
			continue
		}

		if got := buf.String(); got != test.want {
			t.Errorf("encode(%v), got %v, want %v", test.given, got, test.want)
		}
	}
}

func TestDecode(t *testing.T) {
	type Data struct {
		N  *int
		B1 bool
		B2 bool
		M  map[string]int
		S  struct {
			I int
		}
	}

	given := Data{
		N:  nil,
		B1: true,
		B2: false,
		M:  map[string]int{"foo": 0, "bar": 1},
		S:  struct{ I int }{1},
	}

	var buf bytes.Buffer
	err := encode(&buf, reflect.ValueOf(given))
	if err != nil {
		t.Fatalf("error occured while encoding %v. %v", given, err)
	}

	var decoded Data
	err = json.Unmarshal(buf.Bytes(), &decoded)
	if err != nil {
		t.Fatalf("error occured while decoding %v, encoded value of %v. %v", buf.String(), given, err)
	}

	if !reflect.DeepEqual(given, decoded) {
		t.Fatalf("decode(encode(%v)) does not equal to raw value.", given)
	}
}
