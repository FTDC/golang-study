package version

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"gopkg.in/natefinch/npipe.v2"
	"net"
	"os/exec"
)

const (
	channelName  = "top.kikt/go/version"
	getVersion   = "getVersion"
	openUrl      = "openUrl"
	initVpn      = "initVpn"
	connetVpn    = "connetVpn"
	colseConnect = "colseConnect"
)

type VersionPlugin struct{}

type Command struct {
	Fnc     string                   `json:"fnc"`
	Parames []map[string]interface{} `json:"parames"`
}

var _ flutter.Plugin = &VersionPlugin{}

func (VersionPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc(getVersion, getVersionFunc)
	channel.HandleFunc(openUrl, openUrlFunc)
	channel.HandleFunc(initVpn, initVpnFunc)
	channel.HandleFunc(connetVpn, connetVpnFunc)
	channel.HandleFunc(colseConnect, colseConnectFunc)

	//channel := plugin.BasicMessageChannel(messenger, "\\\\.\\pipe\\VPNMainWindow", plugin.StandardMessageCodec{}}
	//channel.HandleFunc(func(_ interface{}) (interface{}, error) { return nil, nil })
	return nil;
}

// 初始化VPN
func initVpnFunc(arguments interface{}) (reply interface{}, err error) {

	argsMap := arguments.(map[interface{}]interface{})

	pac := argsMap["pac"].(string)

	ln, err := npipe.Listen(`\\.\pipe\VPNMainWindow`)

	if err != nil {
		// handle error
	}

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

			fmt.Fprintln(conn, getJsonData("init"))

			r := bufio.NewReader(conn)
			msg, err := r.ReadString('}')
			if err != nil {
				// handle error
				return
			}
			fmt.Println(msg)

			backMsg := jsonToMap(msg)

			fmt.Println(backMsg)
			if (backMsg["fnc"] == "initBack") {
				fmt.Println("init back")

				// 设置pac
				if _, err := fmt.Fprintln(conn, getJsonData("setPacUrl")); err != nil {
					// handle error
				}

				r := bufio.NewReader(conn)
				msg, err := r.ReadString('}')
				if err != nil {
					// handle error
					return
				}
				fmt.Println(msg)

				backMsg := jsonToMap(msg)

				fmt.Println(backMsg)
				if backMsg["fnc"] == "setPacUrlBack" {
					//启动链接
					if _, err := fmt.Fprintln(conn, getJsonData("startXRouteVPN")); err != nil {
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
			}

		}(conn)
	}


	//cmd := exec.Command("XRoute.exe", "")
	//err = cmd.Start()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	return
}

// 链接 VPN
func connetVpnFunc(arguments interface{}) (reply interface{}, err error) {

	return
}

// 关闭VPN
func colseConnectFunc(arguments interface{}) (reply interface{}, err error) {

	return
}

/**
 *  打开系统浏览器，并跳转到指定网页中
 */
func openUrlFunc(arguments interface{}) (reply interface{}, err error) {

	argsMap := arguments.(map[interface{}]interface{})

	url := argsMap["url"].(string)

	cmd := exec.Command("explorer", url)
	//cmd := exec.Command("XRoute.exe", "")
	err = cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	return url, nil;
}

func getVersionFunc(arguments interface{}) (reply interface{}, err error) {

	//cmd := exec.Command("explorer", "https://www.baidu.com")

	//fmt.Println("start server...")
	//listen, err := net.Listen("tcp", "192.168.88.10:8858")
	//if err != nil {
	//	fmt.Println("listen failed, err:", err)
	//	return
	//}
	//for {
	//	conn, err := listen.Accept() //监听是否有连接
	//	if err != nil {
	//		fmt.Println("accept failed, err:", err)
	//		continue
	//	}
	//	go process(conn) //创建goroutine,处理连接
	//}

	return "0.0.1", nil
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
	command := &Command{};
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
		//params := make(map[string]interface{})
		//params["method"] = "aes-256-cfb"
		//params["password"] = "58Ssd2nn95"
		//params["url"] = "120.79.96.245"
		//params["port"] = "8101"
		//params["token"] = "0|0|test34qcPxEJcrE4xVLa41J5"
		//
		//command.Parames = params
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

	return string(data);
}
