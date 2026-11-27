package inventory

import "testing"

func TestAddAndTotal(t *testing.T) {
	s := New()
	s.Add("a", 2)
	s.Add("b", 3)
	if s.Total() != 5 {
		t.Fatalf("got %d", s.Total())
	}
}

func TestRemoveTooMuch(t *testing.T) {
	s := New()
	s.Add("a", 1)
	if err := s.Remove("a", 2); err == nil {
		t.Fatal("expected error")
	}
}
