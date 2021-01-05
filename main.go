package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("Done")
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

type person struct {
	name    string
	age     int
	favFood []string
}

//  시작점
func main() {
	//len, upperName := lenAndUpper("kmk")
	//fmt.Println(len, upperName)

	//repeatMe("mk", "kim")

	//fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//fmt.Println(canIDrink(17))

	//names := []string{"kim", "mk", "00"}
	//names = append(names, "99")
	//fmt.Println(names)

	// map
	//mk := map[string]string{"name": "kmk", "age": "12"}
	//fmt.Println(mk)

	favFood := []string{"kimchi", "ramen"}
	kmk := person{name: "kmk", age: 12, favFood: favFood}
	fmt.Println(kmk.name)
}
