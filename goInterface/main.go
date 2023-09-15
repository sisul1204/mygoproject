package main

import (
	"fmt"
	. "mygoproject/goInterface/model"
)

type Cart struct {
	goods []Goods
}

func (c *Cart) Add(g Goods) {
	c.goods = append(c.goods, g)
}

func (c *Cart) TotalPrice() float64 {
	var totalPrice float64
	for _, g := range c.goods {
		totalPrice += float64(g.GetNum()) * g.GetPrice()
	}
	return totalPrice
}

//func (c *Cart) Add(g Goods) {
//	c.Goods = append(c.Goods, g)
//}
//
//func (c *Cart) TotalPrice() float64 {
//	var totalPrice float64
//	for _, g := range c.Goods {
//		totalPrice += float64(g.GetNum()) * g.GetPrice()
//	}
//	return totalPrice
//}

func main() {
	b1 := Book{"Go从入门到精通", 50, 2}
	b2 := Book{"Go Action", 60, 1}
	p1 := Phone{Brand: "华为", Price: 6999, Num: 1, Discount: 1.0}
	c := &Cart{}
	c.Add(&b1)
	c.Add(&b2)
	c.Add(&p1)
	fmt.Println(c.TotalPrice())

	var goods Goods
	if goods == nil {
		fmt.Println("goods equal nil")
	}

}
