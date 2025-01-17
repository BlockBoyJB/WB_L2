package pattern

import "cmp"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Используется когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта

Плюсы
	1) Быстрая замена стратегии, без необходимости огромных изменений
	2) Инкапсулирование реализации самой стратегии

Минусы
	1) Усложнение кода путем создания разных реализаций одной логики
	2) Необходимость контролировать ("держать в голове") текущий вариант стратегии

Что в примере?
2 разных варианта сортировки: bubble и merge sort
*/

type Strategy[T cmp.Ordered] interface {
	Sort([]T) []T
}

type BubbleStrategy[T cmp.Ordered] struct{}

func (s *BubbleStrategy[T]) Sort(a []T) []T {
	n := len(a)
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		n--
	}
	return a
}

type MergeStrategy[T cmp.Ordered] struct{}

func (s *MergeStrategy[T]) Sort(a []T) []T {
	if len(a) < 2 {
		return a
	}
	first := s.Sort(a[:len(a)/2])
	second := s.Sort(a[len(a)/2:])
	return merge(first, second)
}

func merge[T cmp.Ordered](first, second []T) []T {
	var i, j int

	result := make([]T, 0, len(first)+len(second))

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			result = append(result, first[i])
			i++
		} else {
			result = append(result, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		result = append(result, first[i])
	}
	for ; j < len(second); j++ {
		result = append(result, second[j])
	}
	return result
}

type Context[T cmp.Ordered] struct {
	strategy Strategy[T]
}

func (c *Context[T]) SetStrategy(strategy Strategy[T]) {
	c.strategy = strategy
}

func (c *Context[T]) Sort(a []T) []T {
	return c.strategy.Sort(a)
}
