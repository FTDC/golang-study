package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   []string
}

type Command2 struct {
	Fnc     string  `json:"fnc"`
	Parames []param `json:"parames"`
}

type param struct {
	name, value string
}

func main() {

	//ser := &Server{};
	//ser.ServerName = "local"
	////ser.ServerIP = "192.168.123.1"
	//ser.ServerIP = []string{}
	//
	//data, _ := json.Marshal(ser)
	//fmt.Println(string(data))

	command2 := &Command2{}
	command2.Fnc = "init1"

	var paramArr []param
	//param := make(map[string]string)

	//var param map[string]string;
	//param := t["name"]
	paramArr = append(paramArr, param{"ggg", "werwe"})
	command2.Parames = paramArr

	fmt.Println(command2)
	data2, _ := json.Marshal(command2)
	fmt.Println(string(data2))
}
