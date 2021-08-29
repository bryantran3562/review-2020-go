package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	//fmt.Println("Hi")

	//BT - Create file

	f, error := os.Create("test2.txt")

	if error != nil {
		log.Fatalln("Can not create file...")
	}

	defer f.Close()

	str := "Test\r\n"

	//BT - Converted to slice bytes.
	bs := []byte(str)

	// //BT - Return number of bytes it writes.
	// //Notes: the = instead of :=
	// //_, error = f.Write(bs)

	// bw, error := f.Write(bs)

	// if error != nil {
	// 	log.Fatalln("Could not write file...")
	// }

	// fmt.Printf("Number of bytes write: %d", bw)

	//BT - You can also write to the buffer first and then flush it to the file.

	//BT - Get the pointer to the buffer
	writeBuffer := bufio.NewWriter(f)

	//BT - You can also write a string to it, instead of each bytes.
	for _, item := range bs {
		writeBuffer.WriteByte(item)
	}

	//####################################################################
	//BT - Don't forget to flush. Otherwise, get a blank file.
	//####################################################################
	writeBuffer.Flush()

}
