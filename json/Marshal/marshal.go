package Marshal

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName,string"`
	ServerIP   string `json:"serverIP,omitempty"`
}
type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Guangzhou_Base", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_Base", ServerIP: "127.0.02"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	fmt.Println(string(b))
}
