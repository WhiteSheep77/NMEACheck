// NMEAExe
package main

import (
	"fmt"
)

func main() {
	str := "PCAS06,0"
	var sum int = 0
	var temp int = 0
	for i := 0; i < len(str); i++ {
		temp = (int)(str[i])
		sum = sum ^ temp
	}

	fmt.Printf("0x%x this shouled be 1B\n", sum)

	str2 := "PCAS06,2"
	sum = 0
	temp = 0
	for i := 0; i < len(str2); i++ {
		temp = (int)(str2[i])
		sum = sum ^ temp
	}

	fmt.Printf("and 0x%x\n", sum)
}
