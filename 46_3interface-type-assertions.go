package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"])
	fmt.Println("Apa yh => ", reflect.TypeOf(data["grade"]))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println("Apa yh => ", reflect.TypeOf(data["hobbies"]))
	fmt.Println(data["hobbies"].([]string))
}
