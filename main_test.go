package main

import "testing"

func Test_fuelConsumption(t *testing.T) {
	tests := []struct {
		name string
		needFuelIn100km float64
		leftFuel float64
		want float64
	}{
		// TODO: Add test cases.

		{"need fuel in 100 km is 0", 0, 5, 0},
		{"left fuel is 0", 10,0, 0},
		{"Have Liters is less than 0", 10, -4, 0.},
		{"Both of them is 0", 0, 0 , 0},
		{"both of them more than 0", 5, 10, 200},
	}
	for _, tt := range tests {
		if got := fuelConsumption(tt.needFuelIn100km, tt.leftFuel); got != tt.want {
			t.Error("fuelConsumption()=",tt.name, "got",got, "want", tt.want)
		}
	}
}

