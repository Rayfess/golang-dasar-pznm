package main

import "fmt"

func main() {
	type CreditCard string // Alias, type declaration

	var JohnHas CreditCard = "sf374fa1u"

	var sample string = "2h7fd38wf"
	var sampleCard CreditCard = CreditCard(sample)

	fmt.Println(JohnHas)
	fmt.Println(sampleCard)
}