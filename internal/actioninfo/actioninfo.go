package actioninfo

import (
	"fmt"
)

// DataParser - это интерфейс, который определяет два метода:
// 1. Parse- разбирает входные данные и записывает их в структуру.
// 2. ActionInfo - возвращает строку с информацией об активности.
type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

// Info  принимает слайс строк с данными о тренировках
// парсит их и выводит информацию о тренировке в случае правильных вводных данных,
// иначе - выводит ошибку
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		s, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		fmt.Println(s)
	}
}
