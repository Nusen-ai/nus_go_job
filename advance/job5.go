package advance

import (
	"fmt"
	"math"
)

// 5. 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func Job5TestDemo() {

	// 创建Rectangle实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形 - 宽度: %.1f, 高度: %.1f\n", rect.Width, rect.Height)
	fmt.Printf("  面积: %.2f\n", rect.Area())
	fmt.Printf("  周长: %.2f\n", rect.Perimeter())

	fmt.Println()

	// 创建Circle实例
	circle := Circle{Radius: 4}
	fmt.Printf("圆形 - 半径: %.1f\n", circle.Radius)
	fmt.Printf("  面积: %.2f\n", circle.Area())
	fmt.Printf("  周长: %.2f\n", circle.Perimeter())

}
