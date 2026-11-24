package main

import "testing"

func TestStack(t *testing.T) {
	s := NewStack(2)
	if _, ok := s.Pop(); ok {
		t.Error("pop on empty must underflow")
	}
	if !s.Push(1) || !s.Push(2) {
		t.Fatal("pushes within capacity must succeed")
	}
	if s.Push(3) {
		t.Error("push beyond capacity must overflow")
	}
	if v, ok := s.Top(); !ok || v != 2 {
		t.Errorf("Top = (%d,%v), want 2", v, ok)
	}
	if s.Size() != 2 {
		t.Errorf("Size = %d, want 2", s.Size())
	}
	if v, _ := s.Pop(); v != 2 {
		t.Errorf("Pop = %d, want 2 (LIFO)", v)
	}
	if v, _ := s.Pop(); v != 1 {
		t.Errorf("Pop = %d, want 1", v)
	}
	if _, ok := s.Top(); ok {
		t.Error("Top on empty must be not-ok")
	}
	// Reusable after draining.
	if !s.Push(9) {
		t.Error("push after drain must succeed")
	}
}
