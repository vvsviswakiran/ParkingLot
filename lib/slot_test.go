package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSlot(t *testing.T) {
	t.Run("should not panic when CreateSlot is called", func(t *testing.T) {
		car := CreateCar("AP90GH2345", "White")
		assert.NotPanics(t, func() {
			CreateSlot(123, false, &car)
		})
	})

	t.Run("should panic when slot number passed is less than 1", func(t *testing.T) {
		car := CreateCar("AP90GH2345", "White")
		assert.Panics(t, func() {
			CreateSlot(-1, false, &car)
		})
	})

	t.Run("should not panic when slot is created with no car parked", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateSlot(23, true, nil)
		})
	})
}
