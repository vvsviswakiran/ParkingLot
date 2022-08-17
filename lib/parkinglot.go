package parkinglot

type ParkingLot struct {
	TotalNumberOfSlots int
	Slots              []Slot
}

func CreateParkingLot(TotalSlots int) ParkingLot {
	if TotalSlots < 0 {
		panic("Slots must be positive")
	}
	Slots := make([]Slot, TotalSlots)
	SlotNumber := 1
	for i := range Slots {
		Slots[i] = CreateSlot(SlotNumber, true, nil)
	}
	return ParkingLot{TotalSlots, Slots}
}

func FindNearestEmptySlot(parkingLot ParkingLot) int {
	for i := range parkingLot.Slots {
		if parkingLot.Slots[i].Availability == true {
			return -1
		}
	}
	panic("No empty slots")
}
