package main

import "time"

func unsafeCountdown(position int) {
	println(position)
	time.Sleep(500 * time.Millisecond)
	unsafeCountdown(position - 1)
}

func safeCountdown(position int) {
	println(position)
	if position <= 1 {
		println("Phew, we made it!")
		return
	} else {
		time.Sleep(500 * time.Millisecond)
		safeCountdown(position - 1)
	}
}

func main() {
	safeCountdown(10)
	unsafeCountdown(10)
}
