package main

import (
	"fmt"
	"testing"
)

func TestResolveStatusCode(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		wantErr     bool
		expectedErr error
	}{
		{"200", "200: OK", false, nil},
		{"404", "404: Not Found", false, nil},
		{"500", "500: Internal Server Error", false, nil},
		{"foo", "", true, fmt.Errorf("foo is not a valid input")},
		{"999", "", true, fmt.Errorf("999 is not a valid status code")},
	}

	for _, tc := range tests {
		result, err := ResolveStatusCode(tc.input)

		if result != tc.expected {
			t.Errorf("input %s, expected %s, got %s", tc.input, tc.expected, result)
		}

		if tc.wantErr && (tc.expectedErr.Error() != err.Error()) {
			t.Errorf("input %s, expected %s, got %s", tc.input, tc.expectedErr.Error(), err.Error())
		}
	}
}
