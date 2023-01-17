//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initCar() CarContainer {
	wire.Build(carContainer, car, year, brand, newCarFactory)
	return CarContainer{}
}
