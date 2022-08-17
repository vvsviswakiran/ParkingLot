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