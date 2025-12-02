package main

import "testing"

func TestDialOperateToZero(t *testing.T) {

	tests := []struct {
		name       string
		dial       int
		operation  string
		count      int
		zeroes     int
		wantDial   int
		wantZeroes int
	}{
		{"R within bounds", 50, "R", 141, 0, 91, 1},
		{"L within bounds", 50, "L", 141, 0, 9, 1},
		{"L close bounds", 2, "L", 2, 0, 0, 1},
		{"R close bounds", 99, "R", 2, 0, 1, 1},
		{"Land on zero", 50, "L", 50, 0, 0, 1},
		{"Wrap around R", 50, "R", 1000, 0, 50, 10},
		{"Wrap around L", 50, "L", 1000, 10, 50, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dialOperate(tt.dial, tt.operation, tt.count, &tt.zeroes)

			if got != tt.wantDial {
				t.Errorf("Name = %s, dial = %d, want %d", tt.name, got, tt.wantDial)
			}
			if tt.zeroes != tt.wantZeroes {
				t.Errorf("Name = %s, zeroes = %d, want %d", tt.name, tt.zeroes, tt.wantZeroes)
			}
		})
	}
}
