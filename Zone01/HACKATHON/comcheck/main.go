package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	for _, char := range arg {
		if char == "01" || char == "galaxy" || char == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
