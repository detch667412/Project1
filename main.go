package main

import (
	"fmt"
)

func main() {
	// //(1)FizzBuzz
	// println(fizzbuzz.Fizzbuzz(3))

	// //(2)ReadFile
	// filename := "input.csv"
	// delimiter := '|'
	// records, err := readfile.ReadCSVFile(filename, delimiter)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("CSV Records:")
	// for _, record := range records {
	// 	fmt.Println(record)
	// }
	// // fmt.Printf("%#v", records)

	// type Student struct {
	// 	Id   string
	// 	Name string
	// }

	// temp := []Student{
	// 	{
	// 		Id:   "1",
	// 		Name: "test",
	// 	},
	// 	{
	// 		Id:   "2",
	// 		Name: "test2",
	// 	},
	// 	{
	// 		Id:   "3",
	// 		Name: "test3",
	// 	},
	// }
	// fmt.Println(temp[1])

	// Data to be written to the CSV file

}

// //////////////////////////////
func Add(x, y int) int {
	return x + y
}

func NotEqual(x string) error {
	if x == "error" {
		return fmt.Errorf("incorrect")
	}
	return nil
}
