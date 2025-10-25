//go:build integration
// +build integration

package pacecalc

import (
	"testing"
	"time"
)

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
