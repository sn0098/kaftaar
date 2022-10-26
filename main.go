package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open("companies.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		company := scanner.Text()
		fmt.Println(company)
		exec.Command("open", company).Start()
		fmt.Scanln()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
