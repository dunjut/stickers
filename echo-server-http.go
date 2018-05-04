package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", echoHandler)
	err := http.ListenAndServe(":10001", nil)
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}
}

type echoResp struct {
	Server string `json:"server"`
	ReqURI string `json:"request_uri"`
	Host   string `json:"host"`
}

func echoHandler(w http.ResponseWriter, req *http.Request) {
	resp := echoResp{
		Server: getLocalIP(),
		ReqURI: req.RequestURI,
		Host:   req.Host,
	}
	respdata, err := json.Marshal(resp)
	if err != nil {
		log.Println("json.Marshal:", err)
		w.WriteHeader(500)
		return
	}
	w.Write(respdata)
	log.Println("- OK -")
	return
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "N/A"
	}

	for _, addr := range addrs {
		ip := addr.String()
		if strings.HasPrefix(ip, "10.") {
			return ip
		}
	}

	return "N/A"
}
