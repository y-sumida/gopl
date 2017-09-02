package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC			Celsius = 0
	BoilingC			Celsius = 100
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FtoC(f Fahrenheit) Celsius { return Celsius(f-32) * 5/9}

func (c Celsius) String() string { return fmt.Sprintf("%g℃", c)}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC))
	//fmt.Print("%g\n", boilingF-FreezingC) // invalid operation: boilingF - FreezingC (mismatched types Fahrenheit and Celsius)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	//fmt.Println(c == f) // invalid operation: c == f (mismatched types Celsius and Fahrenheit)
	fmt.Println(c == Celsius(f))

	c = FtoC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}



