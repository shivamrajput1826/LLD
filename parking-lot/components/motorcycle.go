package parkinglot

func NewMotorCycle(licensePlate string) Vehicle {
	return &BaseVehicle{licensePlate: licensePlate, vehicleType: MOTORCYCLE}
}
