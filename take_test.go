package iter

import "testing"

func TestTake(t *testing.T) {
	initialIter := Range(0, 10, 1)
	takeIter := initialIter.Take(5)

	actualTakeCount := takeIter.Count()
	expectedTakeCount := 5

	if actualTakeCount != expectedTakeCount {
		t.Fatalf("got %v, expected %v", actualTakeCount, expectedTakeCount)
	}

	actualInitialCount := initialIter.Count()
	expectedInitialCount := 5

	if actualInitialCount != expectedInitialCount {
		t.Fatalf("got %v, expected %v", actualInitialCount, expectedInitialCount)
	}
}
