package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"Fnc":"init1","Parames":[{"name":"sd","value":"vv"},{"name":"ggg2","value":"werwe2"},{"name":"gggw","value":"werwee"}]}	`
	var dat map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &dat)
	fmt.Println(dat)

	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
	}
}
