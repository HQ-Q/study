package main

import "fmt"

//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
//然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Area 面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64 //半径
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	r := Rectangle{Width: 10, Height: 20}
	fmt.Println(r.Perimeter())
	fmt.Println(r.Area())

	c := Circle{Radius: 10}
	fmt.Println(c.Perimeter())
	fmt.Println(c.Area())
}
