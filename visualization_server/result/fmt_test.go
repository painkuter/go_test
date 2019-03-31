package result

import "testing"

func TestResults(t *testing.T) {
	for i := 0; i < 10; i++ {
		Results(10, 20)
	}
}

func TestResultsf(t *testing.T) {
	for i := 0; i < 10; i++ {
		Resultsf("%d, %d\n", 10, 20)
	}
}

func TestResultsKV(t *testing.T) {
	for i := 0; i < 10; i++ {
		ResultsKV("min", 10, "max", 20)
	}
}
