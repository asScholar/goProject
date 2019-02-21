// test013
package main

import (
	"fmt"
)

func main() {

	/*定义map*/
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	/*map赋值*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"

	/*遍历访问map*/
	for country := range countryCapitalMap {
		fmt.Println(country, countryCapitalMap[country])
	}

	/*元素是否存在于map*/
	capital, ok := countryCapitalMap["美国"]
	fmt.Println(capital, ok)

	capital1, ok1 := countryCapitalMap["France"]
	fmt.Println(capital1, ok1)

	delete(countryCapitalMap, "France")
	capital2, ok2 := countryCapitalMap["France"]
	fmt.Println(capital2, ok2)
}
