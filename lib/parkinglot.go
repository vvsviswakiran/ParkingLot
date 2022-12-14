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

func (parkingLot *ParkingLot) AddCarAndIssueTicket(car *Car) *Ticket {
	freeSlot := FindNearestEmptySlot(*parkingLot)
	parkingLot.Slots[freeSlot-1].CarParked = car
	parkingLot.Slots[freeSlot-1].Availability = false
	NewTicket := CreateTicket(parkingLot.Slots[freeSlot-1].SlotNumber, car)
	return &NewTicket
}

func SlotNumberOfCarWithGivenRegistrationNumber(parkingLot ParkingLot, regNo string) int {
	if len(regNo) != 10 {
		panic("Length of registration number must be 10")
	}
	for i := range parkingLot.Slots {
		if parkingLot.Slots[i].CarParked.RegistrationNumber == regNo {
			return parkingLot.Slots[i].SlotNumber
		}
	}
	panic("Car not found")
}

func (parkingLot *ParkingLot) DeleteTicketAndFreeSlot(ticket *Ticket) {
	if ticket == nil {
		panic("Invalid ticket")
	}
	parkingLot.Slots[ticket.SlotNumber-1].Availability = true
}

func RegistrationNumbersOfCarsWithGivenColor(parkingLot ParkingLot, color string) []string {
	var ListOfRegistrationNumbers []string
	CountOfCarsWithGivenColor := 0
	for i := range parkingLot.Slots {
		if parkingLot.Slots[i].Availability == false && parkingLot.Slots[i].CarParked.Color == color {
			CountOfCarsWithGivenColor += 1
			ListOfRegistrationNumbers = append(ListOfRegistrationNumbers, parkingLot.Slots[i].CarParked.RegistrationNumber)
		}
	}
	if len(ListOfRegistrationNumbers) > 0 {
		return ListOfRegistrationNumbers
	}
	panic("No cars of given color found")
}

func SlotNumberOfAllSlotsHavingACarWithGivenColor(lot ParkingLot, color string) []int {
	var ListOfSlotNumbers []int
	for i := range lot.Slots {
		if lot.Slots[i].Availability == false && lot.Slots[i].CarParked.Color == color {
			ListOfSlotNumbers = append(ListOfSlotNumbers, lot.Slots[i].SlotNumber)
		}
	}
	if len(ListOfSlotNumbers) > 0 {
		return ListOfSlotNumbers
	}
	panic("No cars of given color found")
}
