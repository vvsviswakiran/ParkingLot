package parkinglot

type Slot struct {
	SlotNumber   int
	Availability bool
	CarParked    *Car
}

func CreateSlot(slotNo int, availability bool, car *Car) Slot {
	if slotNo < 1 {
		panic("Slot number must be positive")
	}
	return Slot{slotNo, availability, car}
}
