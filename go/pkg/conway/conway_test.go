package conway

import "testing"

func TestDoSomething(t *testing.T) {
	want := 123
	if got := DoSomething(); got != want {
		t.Errorf("DoSomething() = %d; want %d", got, want)
	}
}
