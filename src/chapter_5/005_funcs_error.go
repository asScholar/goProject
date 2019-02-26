// 005_funcs_error
package main

import (
	"errors"
	"fmt"
)

/*自定义除0错误*/
var errDivisionByZero = errors.New("division by zero")

func div(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	/*正常运算*/
	return dividend / divisor, nil
}

func main() {
	fmt.Println(div(1, 0))
}
