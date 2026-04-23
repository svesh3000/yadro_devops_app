package application

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcAverage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{2.5, 20.1, 14.75, 3.0}, 10.0875},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input=%v", tc.input), func(t *testing.T) {
			t.Parallel()
			result := calcAverage(tc.input)
			require.InDelta(t, tc.expected, result, 1e-9)
		})
	}
}

func TestCalcMedian(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{2.5, 5.0, 12.4, 3.0}, 4.0},
		{[]float64{2.0, 5.5, 4.3, 8.7, 6.2}, 5.5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input=%v", tc.input), func(t *testing.T) {
			t.Parallel()
			result := calcMedian(tc.input)
			require.InDelta(t, tc.expected, result, 1e-9)
		})
	}
}
