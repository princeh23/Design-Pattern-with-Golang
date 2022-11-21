package main

import "fmt"

type Vehicle interface {
	Go()
}

type Car struct {
}

func (r *Car) Go() {
	fmt.Println("use car")
}

type Bicycle struct {
}

func (r *Bicycle) Go() {
	fmt.Println("use Bicycle")
}

type Traveler struct {
	impl Vehicle
}

func (r *Traveler) SetVehicle(i Vehicle) {
	r.impl = i
}

func (r *Traveler) Go() {
	r.impl.Go()
}

func main() {
	traveler := Traveler{}
	traveler.SetVehicle(&Car{})
	traveler.Go()
	traveler.SetVehicle(&Bicycle{})
	traveler.Go()
}
