//go:build unit
// +build unit

package pacecalc

import (
	"math"
	"testing"
	"time"
)

type PaceTestCase struct {
	name            string
	distanceKm      float64
	duration        time.Duration
	expectedPace    float64
	expectedDisplay *PaceDisplay
}

const eps = 0.01

var pc = PaceCalculator{}

// Test pace calculation on various inputs
func TestCalculatePace(t *testing.T) {
	tests := []PaceTestCase{
		{"50 min 10K", 10, time.Minute * 50, 5.0, nil},
		{"3 hour marathon", 42, time.Hour * 3, 4.29, nil},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actualPace := pc.CalculatePace(test.distanceKm, test.duration)
			if math.Abs(actualPace-test.expectedPace) > eps {
				t.Errorf("%s: got %f: want %f", test.name, actualPace, test.expectedPace)
			}
		})
	}
}
