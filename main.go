package main


func main() {
	// STRINGS
	// var name string = "Hayzedd"
	// var color = "blue"
	// var hobby string
	// fmt.Println(name, color, hobby)
	// name = "Azeez"
	// color = "red"
	// hobby = "reading"
	// // cant be used outside of a function + another way of declaring variables
	// AnotherVariable := "random"
	// fmt.Println(name, color, hobby, AnotherVariable)

	// INTS
	// var numOne int = 25
	// var numTwo = 35
	// numThree := 45

	// // BITS AND MEMORY
	// var numFour int8 = 126
	// var numFive uint = 25
	// var numSix uint8 = 255

	// FLOATS
	// var scoreOne float32 = 25.47
	// var scoreTwo float64 = 88888888888888.675
	// scoreThree := 1.3
	// age := 35
	// name := "Hayzedd"
	// // fmt.Println("I am", age, "years old", "and my name is", name)

	// // Formatting strings %_ = format specifiers
	// fmt.Printf("my name is %v and i am %v years old \n", name, age)
	// fmt.Printf("my name is %q and i am %q years old \n", name, age)
	// fmt.Printf("age is of type %T \n", age)
	// fmt.Printf("You scored %0.1f points \n", 222.555)
	// // Save formatted strings(Sprintf)
	// str:= fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	// fmt.Println(str)

	// ARRAYS

	// var arr [3]int = [3]int{23, 24, 25}
	// var arr1 = [5]int{25, 26, 26, 37, 24}
	// names := [2]string{"azeez", "alhameen"}
	// names[1] = "hayzedd"
	// fmt.Println(arr, len(arr1), names, len(names))
	// fmt.Printf("names is of type %T" , names)

	// SLICE(use array under the hood)
	// var SliceArr = []int{23, 45,67}
	// SliceArr[2] = 65
	// SliceArr = append(SliceArr, 455)
	// fmt.Println(SliceArr, len(SliceArr))
	// lastArrElement := SliceArr[len(SliceArr)-1]
	// fmt.Println(lastArrElement)

	// sliceRanges

	// rangeOne := SliceArr[1:3]
	// rangeTwo := SliceArr[1:]
	// rangeThree := SliceArr[:3]
	// rangeThree = append(rangeThree, 4)
	// fmt.Println(rangeOne , rangeTwo , rangeThree)

	// Using Go Packages
	// STRINGS PACKAGE
	// names := "hayzedd alhameen adeyemi"
	// nameCondition := strings.Contains(names , "alhameen")
	// nameReplace :=  strings.ReplaceAll(names , "hayzedd" , "azeez")
	// nameSplit := strings.Split(names, " ")
	// fmt.Println(nameCondition , nameReplace , nameSplit)

	// SORT method

	// scores := []int{23, 25, 42, 27, 29, 30}
	// sort.Ints(scores)
	// index := sort.SearchInts(scores, 27)
	// fmt.Println(scores, index)

	// names := []string{"azeez", "yoshi", "bowser", "mariam"}
	// sort.Strings(names)
	// strIndex := names[2]
	// strSearch := sort.SearchStrings(names, "azeez")
	// fmt.Println(names , strIndex , strSearch)

	// LOOPS
	// x := 0
	// for x < 7 {
	// 	fmt.Println("value of x is :", x)
	// 	x++
	// }
	// for i:= 0; i < 5 ;i++{
	// 	fmt.Println("value of i is:",i)
	// // }
	// names := []string{"azeez", "yoshi", "bowser", "mariam"}
	// // for i:=0;i<len(names);i++{
	// // 	fmt.Println(i, names[i])
	// // }
	// for index , value := range names{
	// 	fmt.Printf("Position %v is %v \n" , index, value)
	// }
	// for _, value := range names{
	// 	fmt.Printf("Value is %v \n" ,value)
	// }

	// Conditional Statements
	// age := 37
	// if age >= 40 {
	// 	fmt.Println("You are eligible")
	// } else if age > 35 {
	// 	fmt.Println("You can go for mayor position")
	// } else {
	// 	fmt.Printf("Comeback later, You are %v, that is young!", age)
	// }

	// names := []string{"azeez", "yoshi", "bowser", "mariam"}
	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("Continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2{
	// 		fmt.Println("Breaking at position", index)
	// 		break
	// 	}
	// 	fmt.Printf("The value is %v at position %v \n", value, index)
	// }

	// FUNCTIONS
	// printScore(5)
	// cycleScore([]int{4, 8, 5, 9, 10}, printScore)
	// area := calcArea(15)
	// fmt.Printf("Area is %v \n", area)
	// fn, sn := getInitials("azeez")
	// fmt.Printf("First initial is %v and second is %v", fn, sn)

	// calcPoints("Main")
	// presentScore()
	// for _, p := range points {
	// 	fmt.Println(p)
	// }
	presentScore()

}

var testScore = [3]int{2, 7, 8}
var testStudents = [3]string{"azeez", "rayo", "ayomide"}

// azeez - AZEEZ - a
// func getInitials(n string) (string, string) {
// 	u := strings.ToUpper(n)
// 	s := strings.Split(u, " ")

// 	var initials []string
// 	for _, v := range s {
// 		initials = append(initials, v[:1])
// 	}
// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}
// 	return initials[0], ""

// }

// func printScore(score int) {
// 	fmt.Printf("Your score is %v \n", score)
// }
// func calcArea(r float64) float64 {
// 	return math.Round(math.Pi * r * r)

// }
// func cycleScore(scoreSlice []int, f func(int)) {
// 	for _, v := range scoreSlice {
// 		f(v)
// 	}
// }
