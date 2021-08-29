package main

import (
	"fmt"
	"io/ioutil"
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
	//     This is used for reading a large file and print out line by line.
	//     Create a NewScanner with passing in a file handle.
	//     It will return a scaner pointer.
	//     Then use that scanner pointer to access member method .Scan to loop over the large file.
	//     Then load it with scanner.Text() function for to read a line by line.
	//     The scanner.Scan() will be failed with there is an EOF or no more lines to read.
	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }

	// fmt.Println("===========================================================================================")
	// fmt.Println("ReadFile Method....")
	// fmt.Println("===========================================================================================")

	// //BT - You can also can use the ioutil package to read the entire file one shot. This is for smaller file.

	// bytes, err := ioutil.ReadFile("interface.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s", bytes)

	// fmt.Println("===========================================================================================")
	// fmt.Println("ReadAll....")
	// fmt.Println("===========================================================================================")

	//BT - You can also can use the ioutil package to read the entire file one shot. This is for smaller file.

	bytes1, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", bytes1)
}
