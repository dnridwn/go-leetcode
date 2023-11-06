// https://leetcode.com/problems/seat-reservation-manager/

package seat_reservation_manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeatManager(t *testing.T) {
	seatManager := Constructor(5)
	assert.Equal(t, 1, seatManager.Reserve()) // All seats are available, so return the lowest numbered seat, which is 1.
	assert.Equal(t, 2, seatManager.Reserve()) // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
	seatManager.Unreserve(2)                  // Unreserve seat 2, so now the available seats are [2,3,4,5].
	assert.Equal(t, 2, seatManager.Reserve()) // The available seats are [2,3,4,5], so return the lowest of them, which is 2.
	assert.Equal(t, 3, seatManager.Reserve()) // The available seats are [3,4,5], so return the lowest of them, which is 3.
	assert.Equal(t, 4, seatManager.Reserve()) // The available seats are [4,5], so return the lowest of them, which is 4.
	assert.Equal(t, 5, seatManager.Reserve()) // The only available seat is seat 5, so return 5.
	seatManager.Unreserve(5)                  // Unreserve seat 5, so now the available seats are [5].
}
