package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("powershell", "-Command",
		"ps | sort -desc cpu | select -first 60;")
	output, err := c.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
