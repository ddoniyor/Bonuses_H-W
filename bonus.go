package main

func main() {

	sales := []int{12_000, 8_000, 15_000, 8_000}
	sumOfBonuses(sales)

}

func sumOfBonuses(sales []int) int {

	const lineOfBound = 10_000
	const percent = 5
	var total int

	for _, value := range sales {
		if value > lineOfBound {
			cash := value - lineOfBound
			bonusCash := cash * percent / 100
			total += bonusCash
		}
	}
	return total
}
