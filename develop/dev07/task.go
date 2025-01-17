package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	fmt.Println("start")
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		//sig(2*time.Hour),
		//sig(5*time.Minute),
		sig(1*time.Second),
		//sig(1*time.Hour),
		//sig(1*time.Minute),
		sig(3*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Важным условием является закрытие общего канала, потому что если этого не сделать, то будет печально (дедлок)
	// Поэтому важно закрыть его, когда все channels закрыты

	wg := sync.WaitGroup{} // Вместо wg можно использовать любой семафор. Но с WaitGroup решение смотрится элегантно =)
	mainCh := make(chan interface{})

	for _, c := range channels {
		wg.Add(1)

		go func(ch <-chan interface{}) {
			defer wg.Done()
			for d := range ch { // логика простая - все каналы теперь пишут в один общий
				mainCh <- d
			}
		}(c)
	}

	// Наблюдатель =)
	go func() {
		wg.Wait() // Канал не закроется, пока есть хоть один producer
		close(mainCh)
	}()

	return mainCh
}
