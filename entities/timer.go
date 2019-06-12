package entities

import (
	"log"
	"sync"
	"time"
)

func SetTimer(wg *sync.WaitGroup, c chan bool, d time.Duration) {

	defer wg.Done()

	t := time.Now()
	timeoutTime := t.Add(d)
	timeRemaining := getTimeRemaining(timeoutTime)

	for timeRemaining.t > 0 {
		//fmt.Printf("Time remaining : Days: %d Hours: %d Minutes: %d Seconds: %d\n",
		//	timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
		timeRemaining = getTimeRemaining(timeoutTime)
		time.Sleep(1 * time.Millisecond)
	}
	log.Print("Returning from timer")

	c <- true
}

type countDownTimer struct {
	t, d, h, m, s int
}

func getTimeRemaining(t time.Time) countDownTimer {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countDownTimer{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
