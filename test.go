package main

import "fmt"

var points = []int{20, 70, 40, 100}

// func calcPoints(s string) {
// 	fmt.Println("Recceiving from another file : ", s)
// }

func presentScore() {
	// for _, s := range testScore {
	// 	fmt.Printf("score is %v \n", s)
	// }
	studentScores := map[string]int{
		"Tayo": 67,
		"Dele": 95,
		"Bayo": 54,
	}
	marketPrice := map[string]int{
		"Pepper":400,
		"Maggi":600,
		"Soap": 500,
	}
	fmt.Println(studentScores)
	fmt.Println(studentScores["Tayo"])
	// Looping through a map
	for k, v := range studentScores {
		fmt.Printf("%v, your score is %v \n", k, v)
	}
	updateMarketPrice(marketPrice)
	fmt.Println(marketPrice)
	name:= "alhameen"
	namePointer := &name
	// fmt.Println("name pointer memory address:" ,namePointer)
	// fmt.Println("value at pointer memory address:" ,*namePointer)
	fmt.Println(name)
	updateName(namePointer)
	fmt.Println(name)
	
}

func updateMarketPrice(m map[string]int){
	m["Pepper"] = 1800
}
func updateName(s *string){
	*s = "hayzedd"
}