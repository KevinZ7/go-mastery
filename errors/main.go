package main

import "fmt"

//Custom error type used to for dividing by zero errors
type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("Cannot divide %f by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, nil
	}

	if dividend == 0 {
		err := divideError{dividend: dividend}
		return 0.0, err
	}

	return divisor / dividend, nil
}

func main() {

	fmt.Println("Should have no error")
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Result if 10/2 is %f \n", res)

	fmt.Println("Should have error")
	res, err = divide(0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}

}
