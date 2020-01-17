package bpool

import "testing"

func TestPool_Retrieve(t *testing.T) {
	p := New()
	buf, ret := p.Retrieve()
	if buf == nil {
		t.Fatal("should not nil")
	}
	ret()
}
