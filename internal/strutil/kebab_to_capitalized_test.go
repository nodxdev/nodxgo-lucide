package strutil

import (
	"testing"

	"github.com/nodxdev/nodxgo-lucide/internal/assert"
)

func TestKebabToCapitalized(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a", "A"},
		{"a-b", "A B"},
		{"a-b-c", "A B C"},
		{"ab-cd", "Ab Cd"},
		{"ab-cdE", "Ab CdE"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := kebabToCapitalized(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
