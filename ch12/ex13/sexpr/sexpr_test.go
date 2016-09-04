package sexpr

import (
	"reflect"
	"testing"
)

type Movie struct {
	Title string
	URL   string `sexpr:"u"`
}

func TestMarshal(t *testing.T) {
	data, err := Marshal(Movie{"hello", "world"})
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	got := string(data)
	want := `((title "hello") (u "world"))`
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestUnmarshal(t *testing.T) {
	give := []byte(`((title "hello") (u "world"))`)
	got := Movie{}
	want := Movie{"hello", "world"}

	err := Unmarshal(give, &got)
	if err != nil {
		t.Fatalf("Unmarshal(%s) failed : %v", string(give), err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %s, got %s", want, got)
	}
}
