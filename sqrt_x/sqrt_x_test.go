package sqrt_x

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	for i := 1; i <= rand.Intn(10000); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, int(math.Floor(math.Sqrt(float64(i)))), mySqrt(i))
		})
	}
}
