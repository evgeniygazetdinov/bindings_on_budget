package lib

import (
	"strconv"
	"fmt"
	"github.com/therecipe/qt/widgets"
)


func equalize_operation(operator string, first_digit int, second_digit int) int {
	var result int
	switch operator {
	case "+":
		result = first_digit + second_digit
	case "-":
		result = first_digit - second_digit
	}
	return result
}

func operation(label *widgets.QLabel, input_text string, operator string) {
	res := label.Text()
	res_int, err := strconv.Atoi(res)
	add_integer, err := strconv.Atoi(input_text)
	fmt.Println(res_int, add_integer)
	res = strconv.Itoa(equalize_operation(operator, res_int, add_integer))
	if err != nil {
		fmt.Println("error")
	}
	label.SetText(res)
}

func OPERATION_INPUT(label *widgets.QLabel, input *widgets.QLineEdit, operator string){		
	if label.Text() == "0" {
		label.SetText(input.Text())
	} else {
		operation(label, input.Text(), operator)
	}

}