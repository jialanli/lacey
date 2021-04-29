package main

import (
	"fmt"
	lacia "github.com/jialanli/lacia/utils"
	"reflect"
	"time"
)

func main() {
	m := make(map[int]int, 3)
	for i := 1; i <= 9; i++ {
		m[i] = i
	}
	fmt.Println(m, len(m))
	arr := [5]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(arr).Len())

	i := 230000000
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))

	arr0 := make([]int, 2, 3)
	fmt.Println(arr0, len(arr0), cap(arr0))
	arr0[1] = 5
	fmt.Println(arr0, len(arr0), cap(arr0))
	arr0 = append(arr0, 8, 9, 76)
	fmt.Println(arr0, len(arr0), cap(arr0))
}

func fc() {
	fmt.Println("hhh", lacia.GetTimeStrOfDayTimeByTs(time.Now().Unix()))
}
