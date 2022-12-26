package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var variable1, operat, variable2 string
	fmt.Printf("Введите выражение:\n")
	fmt.Fscan(os.Stdin, &variable1, &operat, &variable2)
	fmt.Print("Вы ввели: " + variable1 + " " + operat + " " + variable2 + "\n")
	checkRim(variable1, variable2, operat)
}

func checkOp(nf int, nt int, op string, rim bool) {
	switch op {
	case "+":
		var reps = strconv.Itoa(nf + nt)
		if rim {
			fmt.Printf("Разность: " + toRim(nf+nt))
		} else {
			fmt.Printf("Разность: " + reps)
		}

	case "-":
		var reps = strconv.Itoa(nf - nt)
		if rim {
			fmt.Printf("Сумма: " + toRim(nf-nt))
		} else {
			fmt.Printf("Сумма: " + reps)
		}

	case "*":
		var reps = strconv.Itoa(nf * nt)
		if rim {
			fmt.Printf("Умножение: " + toRim(nf*nt))
		} else {
			fmt.Printf("Умножение: " + reps)
		}

	case "/":
		var reps = strconv.Itoa(nf / nt)
		if rim {
			fmt.Printf("Деление: " + toRim(nf/nt))
		} else {
			fmt.Printf("Деление: " + reps)
		}
	}
}

func checkRim(nf string, nt string, op string) {

	var rims = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	var flag = [2]bool{false, false}
	var rim bool = false
	var nums = [2]int{0, 0}

	for i := 0; i < 10; i++ {
		if rims[i] == nf {
			flag[0] = true
			nums[0] = i + 1
			for i := 0; i < 10; i++ {
				if rims[i] == nt {
					flag[1] = true
					nums[1] = i + 1
				}
			}
		}
	}

	if flag[0] == true && flag[1] == true {
		rim = true
		checkOp(nums[0], nums[1], op, rim)
	} else if flag[0] == false && flag[1] == false {
		var err error
		nums[0], err = strconv.Atoi(nf)

		if err != nil {
			fmt.Printf("Не корректный ввод: " + nf)
		}
		nums[1], err = strconv.Atoi(nt)
		if err != nil {
			fmt.Printf("Не корректный ввод: " + nt)
		}

		rim = false
		checkOp(nums[0], nums[1], op, rim)
	} else {
		fmt.Printf("Не верный формат")
	}
}

func toRim(res int) string {

	var rims = [12]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "L", "C"}
	var response = ""
	var su, fifthy, ten, one int
	var minus = false

	if res < 0 {
		minus = true
		res = res * -1
	}

	if res >= 90 && res < 100 {
		res = res - 100
		response = response + "XC"
	} else if res >= 100 {
		res = res - 100
		response = response + "C"
	}

	if res >= 40 && res < 50 {
		res = res - 40
		response = response + "XL"
	} else if res >= 50 && res < 90 {
		res = res - 50
		response = response + "L"
	}
	su = res / 100
	res = res - su*100
	fifthy = res / 50
	res = res - fifthy*50
	ten = res / 10
	res = res - ten*10
	for i := 0; i < ten; i++ {
		response = response + rims[9]
	}

	if res < 10 && res >= 5 {
		if res == 9 {
			response = response + rims[8]
		} else {
			one = res % 5
			response = response + rims[4]
			for i := 0; i < one; i++ {

				response = response + rims[0]
			}
		}
	} else if res == 4 {
		response = response + rims[3]
	} else if res < 4 {
		for i := 0; i < res; i++ {
			response = response + rims[0]
		}
	}

	if minus {
		response = "-" + response
	}
	return response
}
