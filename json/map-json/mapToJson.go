package main

import (
	"encoding/json"
	"fmt"
)

type Command struct {
	Fnc     string                 `json:"fnc"`
	Parames map[string]interface{} `json:"parames"`
}

type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {

	var command Command

	command.Fnc = "init"

	params := make(map[string]interface{})
	params["Url"] = "123123"
	params["Grl2"] = "3f123123"

	command.Parames = params

	data, _ := json.Marshal(command)
	fmt.Println(string(data))

}
