package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Hi")

	//BT - Create file

	f, error := os.Create("test.txt")

	if error != nil {
		log.Fatalln("Can not create file...")
	}

	defer f.Close()

	str := "Test\r\n"

	//BT - Converted to slice bytes.
	bs := []byte(str)

	//BT - Return number of bytes it writes.
	//Notes: the = instead of :=
	//_, error = f.Write(bs)

	bw, error := f.Write(bs)

	if error != nil {
		log.Fatalln("Could not write file...")
	}

	fmt.Printf("Number of bytes write: %d", bw)

}
