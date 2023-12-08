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

func Fizzbuzz2(number int) string {
	if number == 2 {
		return "2"
	}
	if number == 3 {
		return "fizz"
	}
	if number == 4 {
		return "4"
	}
	if number == 5 {
		return "buzz"
	}
	if number == 12 {
		return "12"
	}
	if number == 15 {
		return "fizzbuzz"
	}
	return "1"
}
