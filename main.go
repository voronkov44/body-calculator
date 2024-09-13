package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			//fmt.Println("Не заданы параметры для расчета")
			//continue
			panic("Не заданы параметры для расчета")
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальная масса тела")
	case IMT < 30:
		fmt.Println("У вас избыточная масса тела")
	case IMT < 35:
		fmt.Println("У вас 1-я степень ожирения")
	case IMT < 40:
		fmt.Println("У вас 2-я степень ожирения")
	default:
		fmt.Println("У вас 3-я степень ожирения")
	}
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO PARAMS ERROR")
	}
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчет? (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
