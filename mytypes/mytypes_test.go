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
