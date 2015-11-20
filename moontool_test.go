package main

import "testing"
import "time"

func TestJulianDayNumber(t *testing.T) {
	cases := []struct {
		input           MoonTime
		expected_output float64
	}{
		{
			MoonTime(time.Date(2013, 1, 1, 0, 30, 0, 0, time.UTC)),
			float64(2456293.520833),
		},
	}

	var actual_output float64
	for _, c := range cases {
		actual_output = (c.input).JulianDayNumber()
		if actual_output != c.expected_output {
			t.Errorf("JulianDayNumber(%q) == %q, want %q", c.input, actual_output, c.expected_output)
		}
	}
}
