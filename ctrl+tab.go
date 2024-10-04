package main

import (
	"fmt"

	"math/rand"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	min := 25
	max := 50

	randomNumber := rand.Intn(max-min+1) + min
	var pre_randomNumber int

	for {

		for randomNumber == pre_randomNumber {
			randomNumber = rand.Intn(max-min+1) + min
		}

		pre_randomNumber = randomNumber

		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		kb.SetKeys(keybd_event.VK_TAB)

		kb.HasCTRL(true)

		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
		sleepTime := time.Duration(randomNumber) * time.Second

		start := time.Now()

		// Loop to print countdown every second
		for remaining := sleepTime; remaining > 0; remaining = sleepTime - time.Since(start) {
			fmt.Printf("\rSeconds remaining: %d   ", remaining/time.Second)
			time.Sleep(1 * time.Second)
		}
	}

}
