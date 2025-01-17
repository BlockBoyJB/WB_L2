package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется

Плюсы
	1) Избавляет от множества больших условных операторов машины состояний
	2) Упрощает работу с контекстом состояний

Минусы
	1) Неоправданное усложнение кода при малом количестве состояний (и они схожи в реализации)

*/

type State interface {
	Handle()
}

type StateContext struct {
	state State
}

func NewStateContext() *StateContext {
	return &StateContext{
		state: &ConcreteStateA{},
	}
}

func (c *StateContext) SetState(state State) {
	c.state = state
}

func (c *StateContext) Handle() {
	c.state.Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("state A")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("state B")
}
