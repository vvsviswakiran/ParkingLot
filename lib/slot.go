package parkinglot

type Slot struct {
	SlotNumber   int
	Availability bool
}

func CreateSlot(slotNo int, availability bool) {
	if slotNo < 1 {
		panic("Slot number must be positive")
	}
}
