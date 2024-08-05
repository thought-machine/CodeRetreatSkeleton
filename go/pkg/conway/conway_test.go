package conway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoSomething(t *testing.T) {
	want := 123

	// Pure Go
	if got := DoSomething(); got != want {
		t.Errorf("DoSomething() = %d; want %d", got, want)
	}

	// Using testify
	assert.Equal(t, want, DoSomething())
}
