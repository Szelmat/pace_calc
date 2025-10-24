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

func TestPaceDisplays(t *testing.T) {
	tests := []PaceTestCase{
		{"2 hour HM", 21, time.Hour * 2, 5.72, &PaceDisplay{minutes: 5, seconds: 43}},
		{"15 min 5K", 5, time.Minute * 15, 3, &PaceDisplay{minutes: 3, seconds: 0}},
		{"1 min 400m", 0.4, time.Minute, 2.5, &PaceDisplay{minutes: 2, seconds: 30}},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actualDisplay := pc.GetDisplayPaceParts(pc.CalculatePace(test.distanceKm, test.duration))

			if actualDisplay != *test.expectedDisplay {
				t.Errorf(
					"%s display: got %02d:%02d, want %02d:%02d",
					test.name, actualDisplay.minutes, actualDisplay.seconds,
					test.expectedDisplay.minutes, test.expectedDisplay.seconds)
			}
		})
	}
}
