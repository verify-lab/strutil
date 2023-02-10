package strutil

import "testing"

func TestStripNonPrintable(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"MIX", "b\u00a0e\u200bhind\n", "behind"},
		{"NULL", "\u0000test\n", "test"},
		{"HORIZONTAL TABULATION", "\u0009test\n", "test"},
		{"LINE FEED", "test\u000a", "test"},
		{"FORM FEED", "te\u000cst", "test"},
		{"CARRIAGE RETURN", "te\u000dst", "test"},
		{"NEXT LINE", "te\u0085st", "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripNonPrintable(tt.args); got != tt.want {
				t.Errorf("StripNonPrintableRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}
