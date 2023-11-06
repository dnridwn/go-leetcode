// https://leetcode.com/problems/seat-reservation-manager/

package seat_reservation_manager

type SeatManager struct {
	seats []uint64
}

func Constructor(numSeats int) SeatManager {
	numBits := (numSeats + 63) / 64
	seats := make([]uint64, numBits)
	return SeatManager{seats}
}

func (sm *SeatManager) Reserve() int {
	for i, seatBlock := range sm.seats {
		if seatBlock != ^uint64(0) {
			for j := 0; j < 64; j++ {
				mask := uint64(1) << uint(j)
				if seatBlock&mask == 0 {
					seatNumber := i*64 + j
					sm.seats[i] |= mask
					return seatNumber + 1
				}
			}
		}
	}

	return -1
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	if seatNumber <= 0 || seatNumber > len(sm.seats)*64 {
		return
	}

	seatIndex := seatNumber - 1
	i, j := seatIndex/64, seatIndex%64
	mask := uint64(1) << uint(j)

	sm.seats[i] &^= mask
}
