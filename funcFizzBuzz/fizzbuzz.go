package fizzbuzz

import "strconv"

func Fizzbuzz(number int) string {

	if number%3 == 0 && number%5 == 0 {
		return "fizzbuzz"
	} else if number%3 == 0 {
		return "fizz"
	} else if number%5 == 0 {
		return "buzz"
	}
	result := strconv.Itoa(number)
	return result
}

//1. readfile csv ที่มี header แต่ละคอลัมท์คั่นด้วย |
// output : header แต่ละ

//2. call api POST
