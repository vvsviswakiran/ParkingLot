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
		parkingLot := CreateParkingLot(0)
		assert.Panics(t, func() {
			FindNearestEmptySlot(parkingLot)
		})
	})
}
