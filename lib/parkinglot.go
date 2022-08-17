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
			return parkingLot.Slots[i].SlotNumber
		}
	}
	panic("No empty slots")
}

func (parkingLot *ParkingLot) AddCarAndIssueTicket(car *Car) Ticket {
	freeSlot := FindNearestEmptySlot(*parkingLot)
	parkingLot.Slots[freeSlot-1].CarParked = car
	parkingLot.Slots[freeSlot-1].Availability = false
	NewTicket := CreateTicket(parkingLot.Slots[freeSlot-1].SlotNumber, car)
	return NewTicket
}

func SlotNumberOfCarWithGivenRegistrationNumber(parkingLot ParkingLot, regNo string) int {
	if len(regNo) != 10 {
		panic("Length of registration number must be 10")
	}
	return 0
}
