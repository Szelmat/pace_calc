package pacecalc

import "time"

// Shared test types and helpers used by both unit and integration tests.
type PaceTestCase struct {
	name            string
	distanceKm      float64
	duration        time.Duration
	expectedPace    float64
	expectedDisplay *PaceDisplay
}

// Single shared PaceCalculator instance for tests
var pc = PaceCalculator{}
