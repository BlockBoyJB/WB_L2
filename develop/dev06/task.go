package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	flagFields    = flag.String("f", "", "choose number of fields")
	flagDelimiter = flag.String("d", "\t", "use custom separator")
	flagSeparated = flag.Bool("s", false, "lines only with separator")
)

type flags struct {
	fields    string
	delimiter string
	separated bool
}

func main() {
	flag.Parse()

	f := flags{
		fields:    *flagFields,
		delimiter: *flagDelimiter,
		separated: *flagSeparated,
	}
	cut(f)
}

func cut(f flags) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Println(parse(sc.Text(), f))
	}
}

func parse(line string, f flags) string {
	line = strings.TrimSpace(line)
	fields := strings.Split(line, f.delimiter)

	if len(fields) == 1 && !f.separated {
		return line
	}

	if len(fields) > 1 && f.separated {
		cols, err := parseColumns(f.fields)
		if err != nil {
			return ""
		}
		result := make([]string, 0, len(cols))
		for _, col := range cols {
			if len(fields) < col {
				continue
			}
			result = append(result, fields[col-1])
		}
		return strings.Join(result, " ")
	}
	return ""
}

func parseColumns(fields string) ([]int, error) {
	fields = strings.TrimSpace(fields)
	f := strings.Split(fields, ",") // через "," по умолчанию

	result := make([]int, 0, len(f))

	for _, v := range f {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}
