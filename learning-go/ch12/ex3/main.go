package main

import (
	"fmt"
	"math"
	"sync"
)

func buildMap() map[int]float64 {
	nums := make(map[int]float64)

	for i := range 100_001 {
		nums[i] = math.Sqrt(float64(i))
	}

	return nums
}

var sqrtMapCached func() map[int]float64 = sync.OnceValue(buildMap)

func main() {
	sqrtMap := sqrtMapCached()
	for i := 0; i <= 100_000; i += 1000 {
		fmt.Println(sqrtMap[i])
	}
}
