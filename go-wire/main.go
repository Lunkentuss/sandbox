package main

import "fmt"

func main() {
	carContainer := initCar()
	fmt.Printf("%s\n", string(carContainer.car))
	fmt.Printf("%s\n", string(carContainer.secondCar))
}
