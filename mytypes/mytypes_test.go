package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestToUpper(t *testing.T) {
	var su mytypes.StringUppercaser
	su.Contents.WriteString("am Israel chai!")
	want := "AM ISRAEL CHAI!"
	got := su.ToUpper()
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestSum(t *testing.T) {
	var ib mytypes.IntBuilder
	ib.Contents = []int{3, 5, 7}
	want := 15
	got := ib.Sum()
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestTriple(t *testing.T) {
	t.Parallel()
	var input mytypes.MyInt = 108
	//input := mytypes.MyInt(108) - looks better
	want := mytypes.MyInt(324)
	got := input.Triple()
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
