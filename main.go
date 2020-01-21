package main

func main() {
}
func fuelConsumption(needFuelIn_100km float64, leftFuel float64) float64 {
	if needFuelIn_100km <=0 || leftFuel <= 0{
		return 0
	}
	const km=100
	distanceForLeftFuel:=0.0
	distanceForLeftFuel = (leftFuel*km/needFuelIn_100km)
	return distanceForLeftFuel

}
