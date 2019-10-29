package main

import (
	"encoding/json"
	"fmt"
)

type AlipayRemoteReqStruct struct {
	Ono         string   `json:ono`
	OrderItem   []Item   `json:item`
	OrderRefund []Refund `json:refund`
}
type Item struct {
	Ono string `json:ono`
	Oid int    `json:oid`
}
type Refund struct {
	Ono     string `json:ono`
	Item    int    `json:item`
	Content string `json:content`
	Imgs    string `json:imgs`
	Status  string `json:status`
}

func main() {
	var m AlipayRemoteReqStruct
	m.Ono = "12345"
	m.OrderItem = append(m.OrderItem, Item{Ono: "Shanghai_VPN", Oid: 1})
	m.OrderItem = append(m.OrderItem, Item{Ono: "Beijing_VPN", Oid: 2})
	//for i := 1; i < 6; i++ {
	//	str := []byte("物品")
	//	str = strconv.AppendInt(str, int64(i), 10)
	//	orderi := Item{Ono: string(str), Oid: i}
	//	m.OrderItem = append(m.OrderItem, orderi)
	//}
	bytes, _ := json.Marshal(m)
	fmt.Printf("json:m,%s\n", bytes)
	//var js AlipayRemoteReqStruct
	//err := json.Unmarshal(bytes, &js)
	//if err != nil {
	//	fmt.Printf("format err:%s\n", err.Error())
	//	return
	//}
}
