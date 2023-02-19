package logic

import (
	"math/rand"
	"strconv"
)

func GetRandNum(long int) string {
	var res string
	for i := 0; i < long; i++ {
		num := rand.Intn(10)
		res = res + strconv.Itoa(num)
	}
	return res
}
