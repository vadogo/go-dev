package main

import (
	"fmt"
	"math"
)

const (
	cash          float64 = 23
	oneApplePrice float64 = 5.99
	onePearPrice  float64 = 7
)

func main() {

	fmt.Println(`Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
					Ми маємо 23 грн.
				
				№1:
	Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	`)

	sumPriceFruits := oneApplePrice*9 + onePearPrice*8
	fmt.Println("Для купівлі 9 яблук та 8 груш нам потрібно", sumPriceFruits, "грн.")

	fmt.Println(`№2:
	Скільки груш ми можемо купити?
	`)
	amountPears := math.Floor(cash / onePearPrice)
	fmt.Println("Ми можемо купити", amountPears, "груши")

	fmt.Println(`№3:
	Скільки яблук ми можемо купити?	
	`)
	amountApples := math.Floor(cash / oneApplePrice)
	fmt.Println("Ми можемо купити", amountApples, "яблука")

	fmt.Println(`Задача №4:
	Чи ми можемо купити 2 груші та 2 яблука?
	`)
	allFruits := oneApplePrice*2 + onePearPrice*2
	result := allFruits < cash
	fmt.Println(result, "Ми не можемо купити 2 яблука та 2 груши на наші 23 грн.")
}
