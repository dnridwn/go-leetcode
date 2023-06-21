package implement_queue_using_stacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	obj := Constructor()

	obj.Push(1)
	assert.Equal(t, []int{1}, obj.list)

	obj.Push(2)
	assert.Equal(t, []int{1, 2}, obj.list)

	assert.Equal(t, 1, obj.Peek())

	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, []int{2}, obj.list)
	fmt.Println(obj.list)

	assert.Equal(t, false, obj.Empty())
}
