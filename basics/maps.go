package basics

import (
	"fmt"
	"maps"
)

func main() {

	// var m map[key]valueType
	// mapVariable = make(map[KeyType]valueType)

	// using a Map Literal
	// mapVariable = make(map[KeyType]valueType) {
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exist with key1")
	}

	myMap2 := map[string]int{"a":1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a":1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}
	
	var myMap4 map[string]string

	if myMap4 == nil{
		fmt.Println("The map is initialize to nil value") 
	} else {
		fmt.Println("The map is not initialize to nil value") 
	}

	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	fmt.Println("myMap4 len is", len(myMap4))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}