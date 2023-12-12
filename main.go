package main

import (
	"fmt"
	"project1/fizzbuzz"
	"project1/readfile"
)

func main() {
	//(1)FizzBuzz
	println(fizzbuzz.Fizzbuzz(3))

	//(2)ReadFile
	filename := "input.csv"
	delimiter := '|'
	records, err := readfile.ReadCSVFile(filename, delimiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CSV Records:")
	for _, record := range records {
		fmt.Println(record)
	}
	// fmt.Printf("%#v", records)

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
