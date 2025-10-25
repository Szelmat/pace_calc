// Package that provides methods for calculating running pace.
package pacecalc

import (
	"math"
	"time"
)

type PaceCalculator struct{}

type PaceDisplay struct {
	minutes int32
	seconds int32
}

// Calculate the average pace of the run (min/km) based on distance and time.
func (p *PaceCalculator) CalculatePace(distanceKm float64, result time.Duration) float64 {
	if distanceKm <= 0 {
		return 0
	}

	timeInMins := result.Minutes() * 60

	// Return rounded to two decimals
	return math.Round(timeInMins/distanceKm*100) / 100
}

// Get the minute and second components for the pace instead of a decimal minute value
func (p *PaceCalculator) GetDisplayPaceParts(pace float64) PaceDisplay {
	mins := math.Floor(pace)
	secs := math.Round((pace - mins) * 60)

	// handle rounding overflow
	if secs == 60 {
		mins = mins + 1
		secs = 0
	}

	return PaceDisplay{
		minutes: int32(mins),
		seconds: int32(secs),
	}
}
