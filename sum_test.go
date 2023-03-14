package bazel_for_go

import "testing"

func Test_Sum(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Error("Sum(1, 2) != 3")
	}
}
