package model

type Goods interface {
	GetPrice() float64
	GetNum() int
}

type Book struct {
	Name  string
	Price float64
	Num   int
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) GetNum() int {
	return b.Num
}

type Phone struct {
	Brand    string
	Discount float64
	Price    float64
	Num      int
}

func (p *Phone) GetPrice() float64 {
	return p.Discount * p.Price
}

func (p *Phone) GetNum() int {
	return p.Num
}

//func main() {
//	goods := Phone{"华为", 0.8, 6999, 1}
//	fmt.Println(goods.GetPrice())
//
//}
