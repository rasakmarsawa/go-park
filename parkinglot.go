package main

import "fmt"

type ParkingLot struct {
	capacity int
	slots    []string // index 0 = slot 1
}

func NewParkingLot(cap int) *ParkingLot {
	return &ParkingLot{
		capacity: cap,
		slots:    make([]string, cap),
	}
}

func (p *ParkingLot) Park(car string) {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == "" {
			p.slots[i] = car
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

func (p *ParkingLot) Leave(car string, hours int) {
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] == car {
			charge := calculateCharge(hours)
			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
				car, i+1, charge)
			p.slots[i] = ""
			return
		}
	}
	fmt.Printf("Registration number %s not found\n", car)
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	for i := 0; i < p.capacity; i++ {
		if p.slots[i] != "" {
			fmt.Printf("%d %s\n", i+1, p.slots[i])
		}
	}
}

func calculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + (hours-2)*10
}