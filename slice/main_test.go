package slice_test

import (
	"fmt"
	"github.com/guolei19850528/go_pkgs/slice"
	"testing"
)

func TestRevNoSort(t *testing.T) {
	x := []int{1, 4, 7, 6, 3}
	slice.ReverseNoSort(x)
	fmt.Println(x)
}

func TestReverseNoSortClone(t *testing.T) {
	x := []int{1, 4, 7, 6, 3}
	y := slice.ReverseNoSortClone(x)
	fmt.Println("before x ", x)
	fmt.Println("before y ", y)
	y[0] = 9
	fmt.Println("after x ", x)
	fmt.Println("after y ", y)
}

func TestCartesianProduct(t *testing.T) {
	var x = [][]any{
		[]any{1, 2, 3},
		[]any{4, 5, 6},
		[]any{7, 8, 9},
	}

	data := slice.CartesianProduct(x...)
	fmt.Println(len(data))
	var x1 = [][]any{
		[]any{1, 2},
		[]any{3, 4},
		[]any{5, 6},
	}
	data1 := slice.CartesianProduct(x1...)
	fmt.Println(data1)
}

func TestRand(t *testing.T) {
	x := []int{1, 4, 7, 6, 3}
	fmt.Println(slice.Rand(x))
}
