package design_parking_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParkingSystem(t *testing.T) {
	ps := Constructor(1, 1, 0)
	assert.Equal(t, true, ps.AddCar(1))
	assert.Equal(t, true, ps.AddCar(2))
	assert.Equal(t, false, ps.AddCar(3))
	assert.Equal(t, false, ps.AddCar(1))
}
