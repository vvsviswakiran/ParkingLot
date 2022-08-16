package parkinglot

type Ticket struct {
	SlotNumber int
	CarParked  *Car
}

func CreateTicket(slotNo int, car *Car) {
	if slotNo < 1 {
		panic("Slot number must be positive")
	}
	if car == nil {
		panic("Invalid car object")
	}
}
