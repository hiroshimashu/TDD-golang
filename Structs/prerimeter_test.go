package perimeter

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
}
