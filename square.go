package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (p *Square) End() Point {
	p.start.x += int(p.a)
	p.start.y += int(p.a)
	return p.start
}

func (p *Square) Area() uint {
	return p.a * p.a
}

func (p *Square) Perimeter() uint {
	return p.a * 4
}
