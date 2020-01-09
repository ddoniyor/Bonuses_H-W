package main

func main() {

	sales := []int{12_000, 8_000, 15_000, 8_000}
	sumOfBonuses(sales)

}

func sumOfBonuses(sales []int) int {

	var cash, total int
	const lineOfBound = 10_000
	const percent = 5
	bonusCash := 1

	for i := 0; i < len(sales); i++ {
		if sales[i] > lineOfBound {
			cash = sales[i] - lineOfBound
			bonusCash = cash * percent / 100
			total += bonusCash
		}
	}
	return total
}
