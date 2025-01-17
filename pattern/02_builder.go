package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
	https://refactoring.guru/ru/design-patterns/builder/python/example
*/

/*
Паттерн строитель используется для поэтапного построения сложного продукта

Плюсы
	1) Позволяет создавать различные варианты объектов с тем же процессом создания
	2) Позволяет отделить процесс создания сложного объекта от его использования

Минусы
	1) Избыточность кода
	2) Для простых объектов паттерн избыточен

Что в этом примере?
Реализован "поэтапный" процесс постройки дома. Мы получаем готовый вариант объекта, создавая поэтапно его части
*/

// HouseBuilder абстрактный класс строителя, который описывает требования к постройке
type HouseBuilder interface {
	BuildWalls(walls string)
	BuildDoors(doors string)
	BuildRoof(roof string)
	BuildGarage(garage string)
	Reset()
	GetResult() *House
}

// House Наш итоговый объект, который требует поэтапного строительства
type House struct {
	Walls  string
	Doors  string
	Roof   string
	Garage string
}

// Director сущность, которая управляет процессом построения объекта и определяет порядок
// (в данном примере мы не можем построить крышу без стен, поэтому должен быть порядок)
type Director struct {
	builder HouseBuilder
}

func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

// Build основная функция, которая поэтапно создает сложный объект
func (d *Director) Build() *House {
	d.builder.BuildWalls("white walls")
	d.builder.BuildDoors("pink doors")
	d.builder.BuildGarage("garage for 2 cars")
	d.builder.BuildRoof("roof of tiles")
	return d.builder.GetResult()
}

func (d *Director) ChangeBuilder(builder HouseBuilder) {
	d.builder = builder
}

type ConcreteBuilder struct {
	house *House
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{house: &House{}}
}

func (b *ConcreteBuilder) BuildWalls(walls string) {
	b.house.Walls = walls
}

func (b *ConcreteBuilder) BuildRoof(roof string) {
	b.house.Roof = roof
}

func (b *ConcreteBuilder) BuildGarage(garage string) {
	b.house.Garage = garage
}

func (b *ConcreteBuilder) BuildDoors(doors string) {
	b.house.Doors = doors
}

func (b *ConcreteBuilder) Reset() {
	b.house.Walls = ""
	b.house.Doors = ""
	b.house.Garage = ""
	b.house.Roof = ""
}

func (b *ConcreteBuilder) GetResult() *House {
	return b.house
}
