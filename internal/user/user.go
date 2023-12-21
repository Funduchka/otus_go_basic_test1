package main

import (
	"fmt"

	"github.com/Funduchka/otus_go_basic_test1/pkg/test21"
)

func main() {

	test21.Var2 = "Common"

	fmt.Println(test21.Var2)
	fmt.Println(test21.GetVar1())
	fmt.Println(test21.Var_my)
	//fmt.Println(test21.Var2, test21.GetVar1())
}
