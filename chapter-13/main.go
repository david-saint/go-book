package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		// Handle errors here
		return
	}

	defer file.Close()

	// get the file size
	stat, err := file.Stat()

	if err != nil {
		// Handle errors here
		return
	}

	bs := make([]byte, stat.Size())

	// Read the file
	n1, err := file.Read(bs)

	if err != nil {
		// Handle error here
		return
	}

	/* Alternative way to do it using io/ioutil package */
	// bs, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	return
	// }

	str := string(bs[:n1])
	fmt.Println(str)
}
