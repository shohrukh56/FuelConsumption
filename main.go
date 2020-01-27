package main

func main() {
}
func fuelConsumption(needFuelIn100km float64, leftFuel float64) float64 {
	if needFuelIn100km <=0 || leftFuel <= 0{
		return 0
	}
	const km=100
	return leftFuel*km/needFuelIn100km

}
