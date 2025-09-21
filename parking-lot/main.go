package main

import parkinglot "parking-lot/components"

func main() {
	parkingLot := parkinglot.GetParkingLotInstance()
	parkingLot.AddLevel(parkinglot.NewLevel(1, 100))
	parkingLot.AddLevel(parkinglot.NewLevel(2, 80))

	vehicleFactory := &parkinglot.VehicleFactory{}

	parkingLot.ParkVehicle(vehicleFactory.CreateVehicle("1231", parkinglot.CAR))
	parkingLot.ParkVehicle(vehicleFactory.CreateVehicle("1231414", parkinglot.TRUCK))
	parkingLot.ParkVehicle(vehicleFactory.CreateVehicle("1090i", parkinglot.MOTORCYCLE))

	parkingLot.DisplayAvailability()

	parkingLot.DisplayAvailability()
}
