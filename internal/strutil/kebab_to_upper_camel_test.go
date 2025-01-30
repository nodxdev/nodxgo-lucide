package strutil

import (
	"testing"

	"github.com/nodxdev/nodxgo-lucide/internal/assert"
)

func TestKebabToUpperCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello", "Hello"},
		{"hello-world", "HelloWorld"},
		{"-hello-world", "HelloWorld"},
		{"hello-world-", "HelloWorld"},
		{"-hello-world-", "HelloWorld"},
		{"----", ""},
		{"hello-World", "HelloWorld"},
		{"hello-World-", "HelloWorld"},
		{"-hello-World-", "HelloWorld"},
		{"-hello-World", "HelloWorld"},
		{"hello-World-World", "HelloWorldWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := kebabToUpperCamel(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
