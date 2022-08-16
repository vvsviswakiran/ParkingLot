package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCar(t *testing.T) {
	t.Run("should not panic when CreateCar function is called", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateCar("AP03AL4067", "Green")
		})
	})

	t.Run("should panic when color of car is empty", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateCar("AP03AL2345", "")
		})
	})

	t.Run("should panic when length of registration number is not equal to 10", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateCar("A", "white")
		})
	})
}
