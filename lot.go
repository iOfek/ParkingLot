package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Lot struct {
	size       int
	capacity   int
	regnumber  map[string]Car
	slotnumber map[int]Car
	freeSlots  *IntHeap
}

func makeLot(capacity int) *Lot {
	if capacity <= 0 {
		fmt.Println("number of slots should be positive")
		return nil
	}
	lot := &Lot{size: 0, capacity: capacity, regnumber: make(map[string]Car), slotnumber: make(map[int]Car), freeSlots: &IntHeap{}}
	for i := 1; i <= capacity; i++ {
		heap.Push(lot.freeSlots, i)
	}
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
	return lot
}

func (l *Lot) insertCar(car Car) error {
	if l.size == l.capacity {
		fmt.Println("parking lot full")
		return errors.New("parking lot full")
	}
	car.slot = heap.Pop(l.freeSlots).(int)
	l.slotnumber[car.slot] = car
	l.regnumber[car.reg] = car
	l.size++
	fmt.Printf("Allocated slot number: %d\n", car.slot)
	return nil
}

func (l *Lot) removeCar(reg string) error {
	// check if car is in lot
	c, error := l.getCar(reg)
	if error != nil {
		return error
	}
	//
	delete(l.slotnumber, c.slot)
	delete(l.regnumber, c.reg)
	l.size--
	heap.Push(l.freeSlots, c.slot)
	fmt.Printf("Slot number %d is free\n", c.slot)
	return nil
}

func (l *Lot) getCar(reg string) (Car, error) {
	// check if car is in lot
	_, ok := l.regnumber[reg]
	if !ok {
		fmt.Println("car not in parking lot")
		return Car{}, errors.New("car not in parking lot")
	}
	//
	return l.regnumber[reg], nil
}

func (l *Lot) registration_numbers_for_cars_with_colour(color string) {
	fmt.Printf("Slot No.----Registration No.----Colour\n")
	for _, car := range l.slotnumber {
		if car.color == color {
			fmt.Println(car.reg)
		}
	}
}
func (l *Lot) slot_numbers_for_cars_with_colour(color string) {
	fmt.Printf("Slot No.----Registration No.----Colour\n")
	for _, car := range l.slotnumber {
		if car.color == color {
			fmt.Println(car.slot)
		}
	}
}

func (l *Lot) status() {
	fmt.Printf("Slot No.----Registration No.----Colour\n")
	for _, car := range l.slotnumber { // O(n)
		fmt.Println(car)
	}

}
