package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func (c *Celsius) Add(d Celsius) {
	*c += d
}

func main() {
	var temp float64
	fmt.Print("Введите температуру в Цельсиях: ")
	fmt.Scan(&temp)

	c := Celsius(temp)

	fahrenheit := c.ToFahrenheit()

	fmt.Printf("Температура в Фаренгейтах: %.2f\n", fahrenheit)

	fmt.Printf("Исходная температура: %.2f°C\n", c)
	c.Add(5.0)
	fmt.Printf("После добавления 5°C: %.2f°C\n", c)

	fmt.Printf("Новая температура в Фаренгейтах: %.2f\n", c.ToFahrenheit())
}
