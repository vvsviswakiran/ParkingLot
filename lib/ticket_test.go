package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	t.Run("should not panic when CreateTicket is called", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateTicket(56, &Car{"ER09AL3456", "White"})
		})
	})

	t.Run("should panic when slot number is less than 1", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateTicket(-4, &Car{"ER09AL3456", "White"})
		})
	})

	t.Run("should panic when invalid car object is passed", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateTicket(12, nil)
		})
	})
}
