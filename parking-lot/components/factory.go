package parkinglot

type VehicleFactory struct{}

func (vf *VehicleFactory) CreateVehicle(licensePlate string, vehicleType VehicleType) Vehicle {
	switch vehicleType {
	case CAR:
		return NewCar(licensePlate)
	case MOTORCYCLE:
		return NewMotorCycle(licensePlate)
	case TRUCK:
		return NewTruck(licensePlate)
	default:
		return &BaseVehicle{licensePlate: licensePlate}
	}
}
