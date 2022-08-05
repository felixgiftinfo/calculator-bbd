package calculator

import "errors"

type Calculator struct {
	result int
}

func (cal *Calculator) Clear() {
	cal.result = 0
}

func (cal *Calculator) Press(x int) {
	cal.result = x
}

func (cal *Calculator) Add(x int) {
	cal.result += x
}

func (cal *Calculator) Sub(x int) {
	cal.result -= x
}

func (cal *Calculator) Multiply(x int) (err error) {
	if x == 0 {
		err = errors.New("Zero is prohibitted")
		return
	}
	cal.result *= x
	return
}

func (cal *Calculator) GetResult() int {
	return cal.result
}
