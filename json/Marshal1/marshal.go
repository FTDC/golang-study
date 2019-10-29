package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   []string
}

type Command struct {
	Fnc string `json:fnc`
	//Parames map[string]string `json:"parames"`
	Parames []Param `json:parames`
}

type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {

	//ser := &Server{};
	//ser.ServerName = "local"
	////ser.ServerIP = "192.168.123.1"
	//ser.ServerIP = []string{}
	//
	//data, _ := json.Marshal(ser)
	//fmt.Println(string(data))

	var command Command
	command.Fnc = "init1"

	//var paramArr []Param
	//Param := make(map[string]string)

	//var paramArr []Param;
	//Param := t["name"]
	command.Parames = append(command.Parames, Param{"sd", "vv"})
	//command.Parames = paramArr
	command.Parames = append(command.Parames, Param{"ggg2", "werwe2"})
	command.Parames = append(command.Parames, Param{"gggw", "werwee"})
	//
	//fmt.Println(command.Parames)

	//command.Parames = *command.Parames

	//bytes, _ := json.Marshal(command)
	//fmt.Printf("json:m,%s\n", bytes)

	data2, _ := json.Marshal(command)
	fmt.Println((data2))
	fmt.Println(string(data2))
}
