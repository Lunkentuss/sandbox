package main

type Car string
type Year string
type Brand string

type CarContainer struct {
	car       Car
	secondCar Car
}

func year() Year {
	return Year("1992")
}
func brand() Brand {
	return Brand("volvo")
}

func car(year Year, brand Brand) Car {
	return Car(string(year) + "-" + string(brand))
}

func newCarFactory() func() Car {
	return func() Car {
		return Car("my-car")
	}
}

func carContainer(car Car, carFactory func() Car) CarContainer {
	return CarContainer{car, carFactory()}
}
