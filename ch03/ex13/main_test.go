package ex13

import "testing"

func TestConst(t *testing.T) {
	is1000(t, "KB", KB)
	is1000(t, "MB", MB/KB)
	is1000(t, "GB", GB/MB)
	is1000(t, "TB", TB/GB)
	is1000(t, "PB", PB/TB)
	is1000(t, "EB", EB/PB)
	is1000(t, "ZB", ZB/EB)
	is1000(t, "YB", YB/ZB)
}

func is1000(t *testing.T, name string, given int) {
	if given != 1000 {
		t.Errorf("\nconst %s is invalid", name)
	}
}
