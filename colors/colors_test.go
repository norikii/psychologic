package colors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var nums = []int{1,2,3,4,5}
var yourNums = []int{9,3,0,7,1}
var numsOneElement = []int{3}
var emptyNums []int

func TestCheckColors(t *testing.T) {
	expected := []int{1,1,1,0,0}
	res := CheckColors(yourNums, nums)
	assert.Equal(t, expected, res)
}

func TestSetColors(t *testing.T) {
	fmt.Println(SetColors())
}

func TestCheckNumber_IsPresent(t *testing.T) {
	isPresent := checkNumber(nums, 1)
	assert.True(t, isPresent)
}

func TestCheckNumber_IsNotPresent(t *testing.T) {
	isPresent := checkNumber(nums, 6)
	assert.False(t, isPresent)
}

func TestCheckNumber_EmptyHayStack(t *testing.T) {
	isPresent := checkNumber(emptyNums, 3)
	assert.False(t, isPresent)
}
