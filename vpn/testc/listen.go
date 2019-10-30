package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/natefinch/npipe.v2"
	"net"
)

type Command struct {
	Fnc     string                   `json:"fnc"`
	Parames []map[string]interface{} `json:"parames"`
}

var ln, err = npipe.Listen(`\\.\pipe\VPNMainWindow`)

func main() {
	//ln, err := npipe.Listen(`\\.\pipe\VPNMainWindow`)

	if err != nil {
		// handle error
	}

	initVPN("http://120.79.96.245:8003/routes/pc_d2o.pac")

	//connectVpnFunc(1, "aes-256-cfb", "58Ssd2nn95", "120.79.96.245", "8101", "0|0|test34qcPxEJcrE4xVLa41J5")

}

func jsonToMap(jsonString string) map[string]interface{} {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &dat); err == nil {
		return dat
	}
	return nil
}

//  获取 json 数据
func getJsonData(cmd string) string {
	command := &Command{}
	switch cmd {

	case "init":

		command.Fnc = "init"
		command.Parames = []map[string]interface{}{}
		//command.ServerIP = "192.168.123.1"
	case "startPoliceVPN":
		command.Fnc = "startPoliceVPN"

		value := make(map[string]interface{})
		value["value"] = "aes-256-cfb"
		command.Parames = append(command.Parames, value)

		password := make(map[string]interface{})
		password["value"] = "58Ssd2nn95"
		command.Parames = append(command.Parames, password)

		url := make(map[string]interface{})
		url["value"] = "120.79.96.245"
		command.Parames = append(command.Parames, url)

		port := make(map[string]interface{})
		port["value"] = "8101"
		command.Parames = append(command.Parames, port)

		token := make(map[string]interface{})
		token["value"] = "0|0|test34qcPxEJcrE4xVLa41J5"
		command.Parames = append(command.Parames, token)

	case "setPacUrl":
		command.Fnc = "setPacUrl"

		pac := make(map[string]interface{})
		pac["value"] = "http://120.79.96.245:8003/routes/pc_d2o.pac"
		command.Parames = append(command.Parames, pac)

		//params := make(map[string]interface{})
		//params["pac"] = "http://120.79.96.245:8003/routes/pc_d2o.pac"
		//command.Parames = params

	case "startXRouteVPN":
		command.Fnc = "startXRouteVPN"

		value := make(map[string]interface{})
		value["value"] = "aes-256-cfb"
		command.Parames = append(command.Parames, value)

		password := make(map[string]interface{})
		password["value"] = "58Ssd2nn95"
		command.Parames = append(command.Parames, password)

		url := make(map[string]interface{})
		url["value"] = "120.79.96.245"
		command.Parames = append(command.Parames, url)

		port := make(map[string]interface{})
		port["value"] = "8101"
		command.Parames = append(command.Parames, port)

		token := make(map[string]interface{})
		token["value"] = "0|0|test34qcPxEJcrE4xVLa41J5"
		command.Parames = append(command.Parames, token)

	case "closePoliceVPN":
		command.Fnc = "closePoliceVPN"
		command.Parames = []map[string]interface{}{}

	case "closeXRouteVPN":
		command.Fnc = "closeXRouteVPN"
		command.Parames = []map[string]interface{}{}

	}
	data, _ := json.Marshal(command)

	fmt.Println(string(data))

	return string(data)
}

func initVPN(pacUrl string) {

	// 设置 初始化命令
	command := &Command{}
	command.Fnc = "init"
	command.Parames = []map[string]interface{}{}
	initStr, _ := json.Marshal(command)
	fmt.Println(string(initStr))

	// 设置 PAC 命令
	command.Fnc = "setPacUrl"
	pac := make(map[string]interface{})
	pac["value"] = pacUrl
	command.Parames = append(command.Parames, pac)
	pacStr, _ := json.Marshal(command)

	fmt.Println(string(pacStr))

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		// handle connection like any other net.Conn
		go func(conn net.Conn) {

			//fmt.Println("connect SUCCESS")

			//if _, err := fmt.Fprintln(conn, getJsonData("init")); err != nil {
			//}

			fmt.Fprintln(conn, string(initStr))

			r := bufio.NewReader(conn)
			msg, err := r.ReadString('}')
			if err != nil {
				// handle error
				return
			}
			fmt.Println(msg)

			backMsg := jsonToMap(msg)

			fmt.Println(backMsg)
			if backMsg["fnc"] == "initBack" {
				fmt.Println("init back")

				// 设置pac
				if _, err := fmt.Fprintln(conn, string("pacStr")); err != nil {
					// handle error
				}

				r := bufio.NewReader(conn)
				msg, err := r.ReadString('}')
				if err != nil {
					// handle error
					return
				}
				fmt.Println(msg)
			}

		}(conn)
	}

	//fmt.Println(string(data))
}

func connectVpnFunc(connectType int, valueStr string, passwordStr string, urlStr string, portStr string, tokenStr string) {
	command := &Command{}
	if connectType == 1 {
		command.Fnc = "startXRouteVPN"
	} else {
		command.Fnc = "startPoliceVPN"
	}

	value := make(map[string]interface{})
	value["value"] = valueStr
	command.Parames = append(command.Parames, value)

	password := make(map[string]interface{})
	password["value"] = passwordStr
	command.Parames = append(command.Parames, password)

	url := make(map[string]interface{})
	url["value"] = urlStr
	command.Parames = append(command.Parames, url)

	port := make(map[string]interface{})
	port["value"] = portStr
	command.Parames = append(command.Parames, port)

	token := make(map[string]interface{})
	token["value"] = tokenStr
	command.Parames = append(command.Parames, token)

	connectJson, _ := json.Marshal(command)
	fmt.Println(string(connectJson))

	conn, err := npipe.Dial(`\\.\pipe\VPNMainWindow`)
	if err != nil {
		// handle error
	}
	if _, err := fmt.Fprintln(conn, string(connectJson)); err != nil {
		// handle error
	}
	r := bufio.NewReader(conn)
	msg, err := r.ReadString('\n')
	if err != nil {
		// handle eror
	}
	fmt.Println(msg)

	//for {
	//	conn, err := ln.Accept()
	//	if err != nil {
	//		// handle error
	//		continue
	//	}
	//
	//	// handle connection like any other net.Conn
	//	go func(conn net.Conn) {
	//
	//		//fmt.Println("connect SUCCESS")
	//
	//		//if _, err := fmt.Fprintln(conn, getJsonData("init")); err != nil {
	//		//}
	//
	//		fmt.Fprintln(conn, string(connectJson))
	//
	//		r := bufio.NewReader(conn)
	//		msg, err := r.ReadString('}')
	//		if err != nil {
	//			// handle error
	//			return
	//		}
	//		fmt.Println(msg)
	//
	//		backMsg := jsonToMap(msg)
	//
	//		fmt.Println(backMsg)
	//		if (backMsg["fnc"] == "initBack") {
	//			fmt.Println("init back")
	//
	//			//// 设置pac
	//			//if _, err := fmt.Fprintln(conn, getJsonData("setPacUrl")); err != nil {
	//			//	// handle error
	//			//}
	//			//
	//			//r := bufio.NewReader(conn)
	//			//msg, err := r.ReadString('}')
	//			//if err != nil {
	//			//	// handle error
	//			//	return
	//			//}
	//			//fmt.Println(msg)
	//
	//		}
	//
	//	}(conn)
	//}
}

func closeXRouteVPN(connectType int) {

	command := &Command{}
	var commandName string
	if connectType == 1 {
		commandName = "closePoliceVPN"
	} else {
		commandName = "closeXRouteVPN"
	}

	command.Fnc = commandName
	command.Parames = []map[string]interface{}{}

	data, _ := json.Marshal(command)

	fmt.Println(string(data))

}
