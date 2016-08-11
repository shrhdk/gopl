package tempflag

import (
	"flag"
	"os"
	"testing"
)

func TestDefaultFlag(t *testing.T) {
	os.Args = []string{"command"}
	f := CelsiusFlag("temp1", 0, "the temperature")
	flag.Parse()
	if *f != Celsius(0) {
		t.Errorf("get %v, want %v", *f, Celsius(0))
	}
}

func TestFlag(t *testing.T) {
	os.Args = []string{"command", "-temp2", "0K"}
	f := CelsiusFlag("temp2", 0, "the temperature")
	flag.Parse()
	if *f != AbsoluteZeroC {
		t.Errorf("get %v, want %v", *f, AbsoluteZeroC)
	}
}
