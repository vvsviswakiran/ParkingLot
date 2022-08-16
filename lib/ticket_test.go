package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	t.Run("should not panic when CreateTicket is called", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateTicket(56, Car{"ER09AL3456", "White"})
		})
	})
}
