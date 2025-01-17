package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Применяется когда есть необходимость параметризовать объекты выполняемым действием

Плюсы
	1) Отсутствие зависимости между объектом-исполнителем и объектом, который их вызывает
	2) Простая реализация повтора и отмены операций

Минусы:
	1) Усложненный код программы

*/

type Command interface {
	Execute()
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) Exec() {
	for _, cmd := range i.commands {
		cmd.Execute()
	}
}

type RemoteController struct{}

func (c *RemoteController) Off() {
	fmt.Println("remote controller off")
}

func (c *RemoteController) On() {
	fmt.Println("remote controller on")
}

type RemoteOnCommand struct {
	rc *RemoteController
}

func (r *RemoteOnCommand) Execute() {
	r.rc.On()
}

type RemoteOffCommand struct {
	rc *RemoteController
}

func (r *RemoteOffCommand) Execute() {
	r.rc.Off()
}
