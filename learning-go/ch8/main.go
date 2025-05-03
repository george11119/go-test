package main

import "fmt"

type NumValue interface {
	~int | ~float32 | ~float64
}

func double[T NumValue](val T) T {
	return val * 2
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

func PrintPrintable[P Printable](p P) {
	fmt.Println(p.String())
}

type PrintableInt int

func (p PrintableInt) String() string {
	return fmt.Sprintf("Printable int: %d", p)
}

type PrintableFloat float64

func (p PrintableFloat) String() string {
	return fmt.Sprintf("Printable float: %.2f", p)
}

func main() {
	fmt.Println(double(2))
	fmt.Println(double(2.5))

	var x PrintableInt = 10
	PrintPrintable(x)

	var y PrintableFloat = 52.52
	PrintPrintable(y)
}
