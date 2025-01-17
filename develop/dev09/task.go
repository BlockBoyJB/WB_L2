package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	if len(os.Args) != 2 {
		log.Fatal("wget url does not provided")
	}
	if err := wget(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}

func wget(url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = r.Body.Close() }()

	file, err := os.Create(fmt.Sprintf("wget_%s.html", r.Request.URL.Hostname()))
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if _, err = io.Copy(file, r.Body); err != nil {
		return err
	}
	return nil
}
