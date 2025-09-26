package pacecalc

import (
	"testing"
	"time"
)

var pc = PaceCalculator{}

// Test pace calculation for 50 minute 10K time
func Test50min10kPaceCalc(t *testing.T) {
	want := 5.0
	got := pc.CalculatePace(10, time.Minute*50)

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func Test3HourMarathonPaceCalc(t *testing.T) {
	want := 4.29
	got := pc.CalculatePace(42, time.Hour*3)

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
