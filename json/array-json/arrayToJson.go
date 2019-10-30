package main

import (
	"encoding/json"
	"fmt"
)

type Command struct {
	Fnc     string              `json:"fnc"`
	Parames []map[string]string `json:"parames"`
}

func main() {

	var command Command

	command.Fnc = "init"

	//params := make(map[string]interface{})

	mapValue := make(map[string]string)
	mapValue["value"] = "123123"
	command.Parames = append(command.Parames, mapValue)

	mapValue1 := make(map[string]string)
	mapValue1["value"] = "1231eeeee23"
	command.Parames = append(command.Parames, mapValue1)

	fmt.Println(command.Parames)

	data, _ := json.Marshal(command)
	fmt.Println(string(data))

}
