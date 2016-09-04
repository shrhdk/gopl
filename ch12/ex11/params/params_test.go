package params

import "testing"

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
