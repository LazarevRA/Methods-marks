package personaldata

import "fmt"

// Personal содержит персональные данные о пользователе
type Personal struct {
	Name   string  //Имя пользователя
	Weight float64 //Вес пользователя
	Height float64 //Рост пользователя
}

// Print() выводит персональную информацию о пользователе
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
