package test21

import (
	"fmt"
	_ "strconv"
)

// var (
// 	myInt  int = 50
// 	myInt3     = 43
// )

// var myInt2 = 45
// var myString = "string"

// const (
// 	MY_NAME = "Daria"
// 	MY_AGE  = 28
// )

// const Game = "CS"

func main() {
	// r2, _ := strconv.ParseInt("28", 2, 46)
	// r10, _ := strconv.ParseInt("28", 10, 64)
	// r16, _ := strconv.ParseInt("28", 16, 64)
	// fmt.Println("2 - ", r2)
	// fmt.Println("10 - ", r10)
	// fmt.Println("16 - ", r16)

	// var myInt = 32
	// myInt4 := "32"

	//MY_AGE = 25 cannot assign smth to MY_AGE as it is const

	// fmt.Println("myInt & myInt2", myInt, myInt2)
	// fmt.Println("myIn4", myInt4)
	// fmt.Println("const", MY_NAME, MY_AGE)
	// fmt.Println("Game", Game)

	s := "hello world"            // в двойных кавычках, на одной строке
	s2 := "hello \n world \u9333" // c непечатными символами
	// если нужно включить в строку кавычки или переносы строки
	// - используем обратные кавычки
	s3 := `<div>hello</div>
"cruel"
'world'
`
	d1 := fmt.Sprintf("%s", "It's me")
	d2 := fmt.Sprintf(`%s`, "It's me")
	fmt.Println(s, s2, s3)
	fmt.Println(d1, d2)

}
