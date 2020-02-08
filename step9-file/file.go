package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hi")

	//BT - Open file and check for error
	f, err := os.Open("interface.txt")

	if err != nil {
		log.Fatalln("Could not open file.")
	}

	//BT - This will close the file for us before main exit.
	defer f.Close()

	//BT - Create a NewScanner bufio and pass the file handle to it.
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
