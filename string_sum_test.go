package string_sum

import "testing"

func TestStringSum(t *testing.T) {
	// 3 + 5
	have, _ := StringSum("3+5")
	want := "8"
	if have != want {
		t.Error("incorrect func")
	}
	// 3 - 5
	have, _ = StringSum("3 - 5")
	want = "-2"
	if have != want {
		t.Error("incorrect func")
	}

	// -3 + 5
	have, _ = StringSum("-3 + 5")
	want = "2"
	if have != want {
		t.Error("incorrect func")
	}

	// -3 - 5
	have, _ = StringSum("-3 - 5")
	want = "-8"
	if have != want {
		t.Error("incorrect func")
	}
}
