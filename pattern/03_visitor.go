package pattern

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Посетитель позволяет извлечь родственные операции из классов, составляющих структуру объектов, поместив их в один класс-посетитель.
Если структура объектов является общей для нескольких приложений, то паттерн позволит в каждое приложение включить только нужные операции.

Плюсы
	1) Упрощает добавление операций, работающих со сложными структурами объектов
	2) Объединяет родственные операции в одном классе

Минусы
	1) Паттерн не оправдан, если иерархия элементов часто меняется

Что в примере?
Реализован подсчет площади фигур
*/

type Visitor interface {
	visitForSquare(square *Square)
	visitForCircle(circle *Circle)
	visitForRectangle(rectangle *Rectangle)
	visitForTriangle(triangle *Triangle)
}

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

type Triangle struct {
	Base   float64 // Основание
	Height float64 // Высота (перпендикуляр из противоположной вершины к основанию)
}

type AreaCalculator struct {
	area float64
}

type Figure interface {
	accept(v Visitor)
}

func (c *AreaCalculator) Calculate(f []Figure) float64 {
	for _, figure := range f {
		figure.accept(c)
	}
	return c.area
}

func (c *AreaCalculator) visitForSquare(square *Square) {
	c.area += square.Side * square.Side
}

func (c *AreaCalculator) visitForCircle(circle *Circle) {
	c.area += 3.14 * circle.Radius * circle.Radius
}

func (c *AreaCalculator) visitForRectangle(rectangle *Rectangle) {
	c.area += rectangle.Length * rectangle.Width
}

func (c *AreaCalculator) visitForTriangle(triangle *Triangle) {
	c.area += 0.5 * triangle.Base * triangle.Height
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (t *Triangle) accept(v Visitor) {
	v.visitForTriangle(t)
}
