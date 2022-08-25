package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lot *Lot

func main() {
	// fmt.Println("Hello, world.ws")
	// lot = makeLot(3)
	// lot.insertCar(Car{reg: "KA-01-HH-1234", color: "White"})
	// lot.insertCar(Car{reg: "KA-02-HH-1234", color: "White"})
	// lot.insertCar(Car{reg: "KA-03-HH-1234", color: "White"})
	// lot.insertCar(Car{reg: "KA-02-HH-1234", color: "White"})
	// lot.insertCar(Car{reg: "KA-02-HH-1234", color: "White"})
	// lot.status()
	// lot.removeCar("KA-02-HH-1234")
	// lot.removeCar("KA-01-HH-1234")
	// lot.status()
	// lot.removeCar("KA-02-HH-1234")
	// lot.insertCar(Car{reg: "KA-01-HH-1234", color: "White"})
	// lot.status()
	args := os.Args
	if len(args) == 1 {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			if text == "exit" {
				break
			}
			parse(text)
		}
	} else if len(args) == 2 {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("File does not exist: %s\n", args[1])
		} else {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// fmt.Println(scanner.Text())
				parse(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
		}
		defer file.Close()

	} else {
		fmt.Println("Usage: ./ParkingLot.exe Optional:<inputFile.txt>")
	}

}
func parse(str string) {
	vals := strings.Split(str, " ")
	fmt.Println(vals)
	switch vals[0] {
	case "create_parking_lot":
		if lot != nil {
			fmt.Println("Already create lot.")
		} else {
			capacity, err := strconv.Atoi(vals[1])
			if err != nil {
				fmt.Println("CREATE: Invalid value for number of slots.")
			} else {
				lot = makeLot(capacity)
			}
		}
	case "park":
		lot.insertCar(Car{reg: vals[1], color: vals[2]})
	case "leave":
		lot.removeCar(vals[1])
	case "status":
		lot.status()
	case "registration_numbers_for_cars_with_colour":
		lot.registration_numbers_for_cars_with_colour(vals[1])
	case "slot_numbers_for_cars_with_colour":
		lot.slot_numbers_for_cars_with_colour(vals[1])
	case "slot_number_for_registration_number":
		lot.getCar(vals[1])
	default:
		fmt.Println("ACTION: Invalid action")
	}
}
