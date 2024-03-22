package service

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintSwitch() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random ke-1:", randomizer.Int())
	fmt.Println("random ke-2:", randomizer.Int())
	fmt.Println("random ke-3:", randomizer.Int())

	isEven := false
	if randomizer.Int()%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}

	switch isEven {
	case true:
		fmt.Println("Random Number is Even")
		fmt.Println()
	case false:
		fmt.Println("Random Number is Odd")
		fmt.Println()
	}
}
