package jichu

import (
	"fmt"
)

func IF_Test(a int) {
	if a > 0 {
		a += 100
		fmt.Println(a)
	} else if a == 0 {
		a = 0
		fmt.Println(a)
	} else {
		a -= 100
		fmt.Println(a)
	}
}

func Range_Test() {
	array := [...]int{1, 2, 3, 4, 5, 6}

	for i, x := range array {
		fmt.Printf("array[%d] = %d\n", i, x)

		if i+1 < len(array) {
			array[i+1] += 1
		}
	}

	fmt.Println(array)

	slice := []int{1, 2, 3}
	for i, x := range slice {
		fmt.Printf("slice[%d] = %d\n", i, x)
		if i == 0 {
			slice[1] = 100
		}
		slice = append(slice, x+100)
	}
	fmt.Println(slice)

	maps := map[string]int{"x": 1, "y": 2, "z": 3}
	for k, v := range maps {
		fmt.Printf("maps[%s] = %d\n", k, v)
		delete(maps, k)
	}
	fmt.Println(maps)
}

func Switch_test(a interface{}) {
	switch a {
	case 1:
		fmt.Printf("Int =%d\n", a)
	case 'b', 'c':
		fmt.Printf("String=%c\n", a)
	case true:
		fmt.Printf("Boolean=", a)
		fmt.Println()
		fallthrough
	default:
		fmt.Println("Defult")
	}
}
