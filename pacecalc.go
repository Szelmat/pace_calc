// Package that provides methods for calculating running pace.
package pacecalc

import (
	"math"
	"time"
)

type PaceCalculator struct{}

// Calculate the average pace of the run (min/km) based on distance and time.
func (p *PaceCalculator) CalculatePace(distanceKm float64, result time.Duration) float64 {
	timeInMins := result.Minutes()

	// Return rounded to two decimals
	return math.Round(timeInMins/distanceKm*100) / 100
}
