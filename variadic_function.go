package main

import "fmt"

// variable argument shortcut to calling slice/array but in function
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func negAll(numbers ...int) int  {
// common approach
/* if len(numbers) == 0 { return 0 }
    result := numbers[0]
     for i := 1; i < len(numbers); i++ {
         result -= numbers[i]
     }
     return result 
*/

// modern approach by go
if len(numbers) == 0 {return 0}
result := numbers[0]

for _, number := range numbers[1:] {
	result -= number
}

return result
}

func main() {
	fmt.Println(sumAll(10, 20, 30, 50))
	fmt.Println(negAll(100, 5))

	// slice param, converting from slice to varagrs
	numbers := []int{10,20,30,40}
	fmt.Println(sumAll(numbers...))
}