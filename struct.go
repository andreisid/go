package main

import "fmt"

type car struct {
	gasPedal     uint16 //min 0 max 65535
	breakPedal   uint16
	steringWheel int16
	topSpeedKmh  float64
}

func (c car) mph() float64 {
	return float64(c.topSpeedKmh / 1.60)
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main3() {
	car1 := car{gasPedal: 223, breakPedal: 0, steringWheel: 221, topSpeedKmh: 200.0}
	car2 := car{223, 0, 221, 200.0}
	fmt.Println(car1.gasPedal)
	fmt.Println(car2.gasPedal)
	fmt.Println(car1.mph())
	car1.newTopSpeed(250)
	fmt.Println(car1.topSpeedKmh)
}
