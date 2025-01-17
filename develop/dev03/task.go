package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	flagK = flag.Int("k", 0, "Указание колонки для сортировки")
	flagN = flag.Bool("n", false, "Сортировка по числовому значению")
	flagR = flag.Bool("r", false, "Сортировка в обратном порядке")
	flagU = flag.Bool("u", false, "Не выводить повторяющиеся строки")
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("require input and output files")
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	_ = file.Close()

	f := flags{
		k: *flagK,
		n: *flagN,
		r: *flagR,
		u: *flagU,
	}

	Sort(result, f)
	file, err = os.OpenFile(flag.Arg(1), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range result {
		_, err = file.WriteString(s + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	_ = file.Close()
}

type flags struct {
	k int
	n bool
	r bool
	u bool
}

func Sort(s []string, f flags) {
	if f.u {
		s = removeDuplicates(s)
	}
	sort.Slice(s, func(i, j int) bool {
		row1 := s[i]
		row2 := s[j]

		fields1 := strings.Fields(row1)
		fields2 := strings.Fields(row2)

		if f.k > 0 && f.k <= len(fields1) && f.k <= len(fields2) {
			if f.n {
				value1, err := strconv.Atoi(fields1[f.k])
				if err != nil {
					log.Fatal(err)
				}
				value2, err := strconv.Atoi(fields2[f.k])
				if err != nil {
					log.Fatal(err)
				}
				return value1 < value2
			}
			return fields1[f.k] < fields2[f.k]
		}
		return row1 < row2
	})

	if f.r {
		reverse(s)
	}
}

func removeDuplicates(s []string) (result []string) {
	exist := make(map[string]bool, len(s))

	for _, row := range s {
		if exist[row] {
			continue
		}
		exist[row] = true
		result = append(result, row)
	}
	return result
}

// Разворачиваем слайс методом двух указателей
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
