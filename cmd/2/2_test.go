package main

import "testing"

func TestCheckInvalidIdV2(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{"Simple number", 121212, true},
		{"False number", 128212, false},
		{"Very simple", 111, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInvalidIdV2(tt.num)

			if got != tt.expected {
				t.Errorf("Name = %s, result = %t, expected = %t", tt.name, got, tt.expected)
			}
		})
	}

}

func TestCountInvalidIds(t *testing.T) {
	tests := []struct {
		name       string
		first      int
		last       int
		expected   int
		expectedV2 int
	}{
		{"Simple range", 10, 20, 1, 1},
		{"R2", 95, 115, 1, 2},
		{"R3", 998, 1012, 1, 2},
		{"R4", 222220, 222224, 1, 1},
		{"R5", 1188511880, 1188511890, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idsAddition := 0
			invalidIdsV2 := 0
			got := countInvalidIds(tt.first, tt.last, &idsAddition, &invalidIdsV2, &idsAddition)

			if got != tt.expected {
				t.Errorf("V1 Name = %s, result = %d, expected %d", tt.name, got, tt.expected)
			}

			if invalidIdsV2 != tt.expectedV2 {
				t.Errorf("V2 Name = %s, result = %d, expected %d", tt.name, invalidIdsV2, tt.expectedV2)
			}
		})
	}
}
