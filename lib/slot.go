package parkinglot

type Slot struct {
	SlotNumber   int
	Availability bool
}

func CreateSlot(slotNo int, availability bool) Slot {
	if slotNo < 1 {
		panic("Slot number must be positive")
	}
	return Slot{slotNo, availability}
}
