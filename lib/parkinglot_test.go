package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	t.Run("should not panic when creating a parking lot", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateParkingLot(12)
		})
	})

	t.Run("should panic when number of slots is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateParkingLot(-2)
		})
	})
}

func TestFindNearestEmptySlot(t *testing.T) {
	t.Run("should panic when there are no slots available", func(t *testing.T) {
		assert.Panics(t, func() {
			FindNearestEmptySlot(CreateParkingLot(0))
		})
	})

	t.Run("should not panic when there are free slots", func(t *testing.T) {
		assert.NotPanics(t, func() {
			FindNearestEmptySlot(CreateParkingLot(1))
		})
	})

	t.Run("should return a valid slot number if slots are available", func(t *testing.T) {
		assert.Less(t, 0, FindNearestEmptySlot(CreateParkingLot(2)))
	})
}

func TestAddCarAndIssueTicket(t *testing.T) {
	t.Run("should not panic if free slots are available", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)
		car := CreateCar("AP90GH2345", "White")
		assert.NotPanics(t, func() {
			parkingLot.AddCarAndIssueTicket(&car)
		})
	})
}

func TestSlotNumberOfCarWithGivenRegistrationNumber(t *testing.T) {
	t.Run("should panic if registration number is not of length 10", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)
		car := CreateCar("AP90GH2345", "White")
		parkingLot.AddCarAndIssueTicket(&car)
		assert.Panics(t, func() {
			SlotNumberOfCarWithGivenRegistrationNumber(parkingLot, "")
		})
	})

	t.Run("should return a valid slot number if car is parked in the parking lot", func(t *testing.T) {
		parkingLot := CreateParkingLot(1)
		car := CreateCar("AP90GH2345", "White")
		parkingLot.AddCarAndIssueTicket(&car)
		assert.Less(t, 0, SlotNumberOfCarWithGivenRegistrationNumber(parkingLot, "AP90GH2345"))
	})
}
