package connect

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/natefinch/npipe.v2"
)

type Command struct {
	Fnc     string                   `json:"fnc"`
	Parames []map[string]interface{} `json:"parames"`
}

func main() {
	connectVpnFunc(1, "aes-256-cfb", "58Ssd2nn95", "120.79.96.245", "8101", "0|0|test34qcPxEJcrE4xVLa41J5")
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
