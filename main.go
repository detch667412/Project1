package main

import "fmt"

func main() {

}

func Add(x, y int) int {
	return x + y
}

func NotEqual(x string) error {
	if x == "error" {
		return fmt.Errorf("incorrect")
	}
	return nil
}
