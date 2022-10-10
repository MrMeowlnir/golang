package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// READ FILE
	readdata, err := ioutil.ReadFile("first.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readdata))

	// WRITE TO FILE
	writedata := []byte("Hello, World!")
	ioutil.WriteFile("second.txt", writedata, 0777) // 0 - FILE, 7 - READ, 7 - WRITE, 7 - EXECUTE ALL (GOD MODE)

	// APPEND TO FILE
	f, err := os.OpenFile("third.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700) // 0 - FILE, 7 - user only RWX
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(" World!")
	if err != nil {
		panic(err)
	}

}
