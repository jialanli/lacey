package lacey

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T){
	arr := []int{1,2,3}
	fmt.Println(arr,len(arr),cap(arr))
	arr = arr[:0]
	fmt.Println(arr,len(arr),cap(arr))
}