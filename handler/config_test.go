package handler

import "testing"

func Test_safeStripe(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{key: "sk_test_a1b2c3d4e5c6g7", want: "sk_test_xxxxxxxxxxxxxx"},
		{key: "whsec_a1b2c3d4e5c6g7", want: "whsec_xxxxxxxxxxxxxx"},
		{key: "a1b2c3d4e5c6g7", want: "xxxxxxxxxxxxxx"},
		{key: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := safeStripe(tt.key); got != tt.want {
				t.Errorf("safeStripe(%v) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func Test_safeString(t *testing.T) {
	tests := []struct {
		key  string
		want string
	}{
		{key: "sk_test_a1b2c3d4e5c6g7", want: "xxxxxxxxxxxxxxxxxxxxxx"},
		{key: "whsec_a1b2c3d4e5c6g7", want: "xxxxxxxxxxxxxxxxxxxx"},
		{key: "a1b2c3d4e5c6g7", want: "xxxxxxxxxxxxxx"},
		{key: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := safeString(tt.key); got != tt.want {
				t.Errorf("safeString(%v) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}
