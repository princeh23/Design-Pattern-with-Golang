package main

import "fmt"

type stationeryInterface interface {
	create()
}

type stationeryPencil struct{}

func (*stationeryPencil) create() {
	fmt.Println("pencil create successful")
}

type stationeryPen struct{}

func (*stationeryPen) create() {
	fmt.Println("pen create successful")
}

type stationeryFactory struct{}

func (s stationeryFactory) generate(name string) stationeryInterface {
	switch name {
	case "pencil":
		return new(stationeryPencil)
	case "pen":
		return new(stationeryPen)
	default:
		return nil
	}
}

func main() {
	factory := stationeryFactory{}
	var stationeryType string

	stationeryType = "pencil"
	pencil := factory.generate(stationeryType)
	pencil.create()

	stationeryType = "pen"
	pen := factory.generate(stationeryType)
	pen.create()
}
