package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	flagAfter   = flag.Int("A", 0, "Show +N lines after match")
	flagBefore  = flag.Int("B", 0, "Show +N lines before match")
	flagContext = flag.Int("C", 0, "Show +-N lines around match")
	flagCount   = flag.Bool("count", false, "Show count matched lines")
	flagIgnore  = flag.Bool("i", false, "Ignore the case")
	flagInvert  = flag.Bool("v", false, "Exclude instead of match")
	flagFixed   = flag.Bool("F", false, "Exact match with a string, not a pattern")
	flagNum     = flag.Bool("n", false, "Show line number")
)

type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Usage: my-grep [FLAGS] PATTERN [INPUT_FILE]")
	}
	pattern := flag.Arg(0)

	var lines []string

	if len(flag.Args()) == 2 {
		if err := readFile(flag.Arg(1), &lines); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := readStdin(&lines); err != nil {
			log.Fatal(err)
		}
	}

	f := flags{
		after:      *flagAfter,
		before:     *flagBefore,
		context:    *flagContext,
		count:      *flagCount,
		ignoreCase: *flagIgnore,
		invert:     *flagInvert,
		fixed:      *flagFixed,
		lineNum:    *flagNum,
	}
	grep(lines, pattern, f)
}

func readStdin(buf *[]string) error {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		*buf = append(*buf, sc.Text())
	}
	return sc.Err()
}

func readFile(path string, buf *[]string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		*buf = append(*buf, sc.Text())
	}
	_ = file.Close()
	return sc.Err()
}

func grep(lines []string, pattern string, f flags) {
	var reg *regexp.Regexp

	// если флаг fixed, то нам не нужно использовать регулярное выражение
	if f.fixed {
		if f.ignoreCase {
			pattern = strings.ToLower(pattern) // приводим к нижнему регистру для флага ignore-case + потом будем сравнивать ToLower строки
		}
	} else {
		p := pattern
		if f.ignoreCase {
			p = "(?i)" + pattern
		}
		var err error
		reg, err = regexp.Compile(p)
		if err != nil {
			log.Fatal(err)
		}
	}
	matches := make([]bool, len(lines))
	for i, line := range lines {
		var match bool
		if f.fixed {
			if f.ignoreCase {
				match = strings.ToLower(line) == pattern
			} else {
				match = line == pattern
			}
		} else {
			match = reg.MatchString(line)
		}

		if f.invert {
			match = !match
		}
		matches[i] = match
	}

	if f.count {
		var k int
		for _, m := range matches {
			if m {
				k++
			}
		}
		fmt.Println("matched lines count:", k)
		return
	}

	for i, match := range matches {
		if !match {
			continue
		}

		start := max(0, i-f.before)
		end := min(len(lines), i+f.after+1)
		if f.context != 0 {
			start = max(0, i-f.context)
			end = min(len(lines), i+f.context+1)
		}

		for j := start; j < end; j++ {
			if f.lineNum {
				fmt.Printf("%d:%s\n", j+1, lines[j])
			} else {
				_, _ = fmt.Fprintln(os.Stdout, lines[j])
			}
		}
	}
}
