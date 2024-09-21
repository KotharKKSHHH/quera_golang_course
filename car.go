package main

import (
	"fmt"
)

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	car := &Car{speed: speed, battery: battery}
	return car
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {

	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	moreCharge := minutes / 2
	charge := moreCharge + car.battery
	if charge > 100 {
		charge = 100
	}
	car.battery = charge

}
func TryFinish(car *Car, distance int) string {
	more := distance / 2
	leftover := car.battery - more
	var time float32
	if leftover > -1 {
		car.battery = leftover
		time = float32(distance) / float32(car.speed)
		return fmt.Sprintf("%.2f", time)
	} else {
		car.battery = 0
		return ""
	}

}
