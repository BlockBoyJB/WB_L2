package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> — смена директории (в качестве аргумента могут быть то-то и то);

- pwd — показать путь до текущего каталога;

- echo <args> — вывод аргумента в STDOUT;

- kill <args> — «убить» процесс, переданный в качесте аргумента (пример: такой-то пример);

- ps — выводит общую информацию по запущенным процессам в формате такой-то формат.

Так же требуется поддерживать функционал fork/exec-команд.

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	sc := bufio.NewReader(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("pwd error:", err)
		}

		fmt.Printf("$ %s>", dir)
		input, err := sc.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		args := strings.Fields(input)
		executeCommand(args)
	}
}

func executeCommand(args []string) {
	switch args[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		if len(args) < 2 {
			fmt.Println("cd: missing argument")
			return
		}
		if err := os.Chdir(args[1]); err != nil {
			fmt.Println("cd error:", err)
		}

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("error find pwd", err)
		}
		fmt.Println(dir)

	case "echo":
		fmt.Println(strings.Join(args[1:], " "))

	case "kill":
		if len(args) < 2 {
			fmt.Println("kill: missing argument")
			return
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("kill: extract pid error:", err)
			return
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("kill: find process by pid %d error: %s\n", pid, err)
			return
		}
		if err = proc.Kill(); err != nil {
			fmt.Printf("kill process %d error: %s\n", pid, err)
			return
		}

	case "ps":
		cmd := exec.Command("tasklist") // Этот ваш виндовс просто имба...
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("ps error:", err)
			return
		}

	case "ls":
		cmd := exec.Command("powershell.exe", "dir")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("ls error:", err)
			return
		}

	default:
		cmd := exec.Command("powershell.exe", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("exec command %s error: %s\n", args[0], err)
			return
		}
	}
}
