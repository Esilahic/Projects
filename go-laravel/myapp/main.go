package main

import (
	"fmt"

	"github.com/Esilahic/Projects/go-laravel/celeritas"
)

func main() {

	result := celeritas.TestFuncn(1, 1)
	fmt.Println(result)

	result = celeritas.TestFunc2(2, 1)
	fmt.Println(result)
}
