package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSlot(t *testing.T) {
	t.Run("should not panic when CreateSlot is called", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateSlot(123, true)
		})
	})

	t.Run("should panic when slot number passed is less than 1", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateSlot(-1, true)
		})
	})
}
