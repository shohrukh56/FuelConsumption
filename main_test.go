package main

import "testing"

func Test_fuelConsumption(t *testing.T) {
	type args struct{
		needFuelIn_100km float64
		leftFuel float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"need fuel in 100 km is 0", args{0, 5}, 0},
		{"left fuel is 0", args{10,0}, 0},
		{"Have Liters is less than 0", args{10, -5}, 0.},
		{"Both of them is 0", args{0, 0 }, 0},
		{"both of them more than 0", args{5, 10}, 200},
	}
	for _, tt := range tests {
		if got := fuelConsumption(tt.args.needFuelIn_100km, tt.args.leftFuel); got != tt.want {
			t.Error("fuelConsumprion()=",tt.name, "got",got, "want", tt.want)
		}
	}
}

