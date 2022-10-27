package main

import "fmt"

type printer interface {
	printContent()
}

type hp struct {
}

func (p *hp) printContent() {
	fmt.Println("hp printer print xxx")
}

type epson struct {
}

func (p *epson) printContent() {
	fmt.Println("epson printer print xxx")
}

type computer interface {
	print()
}

type mac struct {
	myPrinter printer
}

func (m *mac) print() {
	fmt.Println("mac prepare to print...")
	m.myPrinter.printContent()
}

type windows struct {
	myPrinter printer
}

func (w *windows) print() {
	fmt.Println("windows prepare to print...")
	w.myPrinter.printContent()
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	w := windows{}
	w.myPrinter = hpPrinter
	w.print()
	w.myPrinter = epsonPrinter
	w.print()

	m := mac{}
	m.myPrinter = hpPrinter
	m.print()
	m.myPrinter = epsonPrinter
	m.print()
}
