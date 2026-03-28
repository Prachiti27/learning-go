package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// creating + writing to file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = file.WriteString("Hello World!\n")
	if err != nil {
		fmt.Println("Error:", err)
		file.Close()
		return
	}
	file.Close() // close AFTER writing

	// reading full file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))

	// read line by line
	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	file.Close() // close AFTER reading

	// appending to file
	file, err = os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file.WriteString("Appending to file\n")
	file.Close() // close AFTER append

	// deleting file
	err = os.Remove("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}