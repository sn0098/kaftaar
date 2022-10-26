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
	counter := 1
	for scanner.Scan() {
		company := scanner.Text()
		fmt.Println(company)
		exec.Command("open", "https://www.google.com/search?q="+company).Start()
		if counter == 10 {
			fmt.Scanln()
			counter = 0
		}
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
