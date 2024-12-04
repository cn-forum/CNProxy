package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"strings"
	"time"
)

type ProxyConfig struct {
	LocalPort  string
	Targets    []string
	AllowedIPs []string
}

func main() {
	// Hier könnt ihr eure localPorts hinzufügen.
	// Ihr könnt ruhig von server1:30100 auf server2:30200 und dann wieder auf server1:30300
	// Spielt so viel rum wie ihr wollt, allowed NUR eure Proxy IP's und schaltet davor einmal balooProxy

	config := []ProxyConfig{
		{
			LocalPort: "80",
			Targets:   []string{"1.1.1.1:3100",},
			AllowedIPs: []string{
				"*",
			},
		},
		{
			LocalPort: "30100",
			Targets:   []string{"1.1.1.2:3200", "1.1.1.3:3300", "1.1.1.4:3400",},
			AllowedIPs: []string{
				"1.1.1.1",
				"1.1.1.2",
				"1.1.1.3",
				"1.1.1.4",
			},
		},
		{
			LocalPort: "30500",
			Targets:   []string{"1.1.1.2:3900", "1.1.1.3:3500", "1.1.1.4:3700",},
			AllowedIPs: []string{
				"1.1.1.1",
				"1.1.1.2",
				"1.1.1.3",
				"1.1.1.4",
			},
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, proxy := range config {
		go startProxy(proxy)
	}

	select {}
}

func startProxy(proxy ProxyConfig) {
	listener, err := net.Listen("tcp", ":"+proxy.LocalPort)
	if err != nil {
		fmt.Printf("Error starting proxy on port %s: %v\n", proxy.LocalPort, err)
		return
	}
	defer listener.Close()
	currentTime := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("\033[48;5;15m\033[38;2;0;0;0m%s \033[0m\033[38;2;168;50;50m •\033[0m Running on Port: \033[38;2;50;168;82m%s\033[0m\n", currentTime, proxy.LocalPort)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		clientIP := strings.Split(clientConn.RemoteAddr().String(), ":")[0]
		if !isIPAllowed(clientIP, proxy.AllowedIPs) {
			fmt.Printf("\033[48;5;15m\033[38;2;0;0;0m%s \033[0m\033[38;2;168;50;50m •\033[0m %s is not whitelisted!\033[0m\n", currentTime, clientIP)
			clientConn.Close()
			continue
		}

		go handleConnection(clientConn, proxy.Targets)
	}
}

func handleConnection(clientConn net.Conn, targets []string) {
	defer clientConn.Close()

	targetAddress := targets[rand.Intn(len(targets))]

	targetConn, err := net.Dial("tcp", targetAddress)
	if err != nil {
		return
	}
	defer targetConn.Close()

	go io.Copy(targetConn, clientConn)
	io.Copy(clientConn, targetConn)
}

func isIPAllowed(clientIP string, allowedIPs []string) bool {
	for _, ip := range allowedIPs {
		if ip == "*" {
			return true
		}
	}
	for _, ip := range allowedIPs {
		if clientIP == ip {
			return true
		}
	}
	return false
}
