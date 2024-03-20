package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("El puerto %d esta abierto\n", i)
	}
}
