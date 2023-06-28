package design_parking_system

type ParkingSystem struct {
	slots map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slots: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	if ps.slots[carType] <= 0 {
		return false
	}
	ps.slots[carType] -= 1
	return true
}
