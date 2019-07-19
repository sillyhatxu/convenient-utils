package test1

import (
	"fmt"
	"strconv"
)

func main() {
	var groupNumber int64 = 4
	plainDeviceId := "12b53d07ee4f46d8bb14599dd9b247d0"
	randomNum, _ := strconv.ParseInt(plainDeviceId[:30], 16, 64)
	modResult := int64(randomNum) % int64(4)
	if int64(modResult) == groupNumber-1 {
		fmt.Println(true)
	}
	fmt.Println(false)
}
