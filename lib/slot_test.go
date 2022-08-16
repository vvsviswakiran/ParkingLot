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
}
