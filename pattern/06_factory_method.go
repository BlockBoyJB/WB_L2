package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Применяется когда система должна оставаться легко расширяемой путем добавления объектов новых типов

Плюсы
	1) Упрощает добавление новых объектов в программу
	2) Избавляет класс от привязки к конкретным объектам

Минусы
	1) Может привести к созданию больших параллельных иерархий классов,
так как для каждого класса продукта надо создать свой подкласс создателя

Что в примере?
Создание объекта для работы с разными типами файлов объединено одним "билдером"
*/

type fileType uint8

const (
	TxtFile fileType = iota
	JsonFile
	DocFile
)

type Creator interface {
	CreateProduct(ft fileType) Product
}

type Product interface {
	Use()
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) CreateProduct(ft fileType) (Product, error) {
	switch ft {
	case TxtFile:
		return &TxtRepo{}, nil
	case JsonFile:
		return &JSONRepo{}, nil
	case DocFile:
		return &DocRepo{}, nil

	default:
		return nil, errors.New("unknown type of file")
	}
}

type TxtRepo struct{}

func (r *TxtRepo) Use() {
	fmt.Println("saving file into txt format")
}

type JSONRepo struct{}

func (r *JSONRepo) Use() {
	fmt.Println("saving file into json format")
}

type DocRepo struct{}

func (r *DocRepo) Use() {
	fmt.Println("saving file into doc format")
}
