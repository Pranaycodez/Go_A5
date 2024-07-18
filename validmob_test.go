// This file is used to test the IsValidMobileNumber function in the validmobileno package.
package validmobileno

// imported testing package
import "testing"

// TestIsValidMobileNumber tests the IsValidMobileNumber function.
func TestIsValidMobileNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"+14372992532", true},       // Valid number with country code
		{"9494021211", true},         // Valid number without country code
		{"+1234", false},             // Invalid, too short
		{"12345678901234567", false}, // Invalid, too long
		{"+44 20 7946 0958", false},  // Invalid format
		{"", false},                  // Invalid, No input
	}

	for _, test := range tests {
		result := IsValidMobileNumber(test.input)
		if result != test.expected {
			t.Errorf("IsValidMobileNumber(%q) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
