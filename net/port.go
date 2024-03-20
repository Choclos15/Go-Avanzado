package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

// go run net/port.go --site=scanme.webscantest.com
var site = flag.String("site", "scanme.nmap.org", "url para escanear")

func main() {
	/*
		for i := 0; i < 100; i++ {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
			if err != nil {
				continue
			}
			conn.Close()
			fmt.Printf("El puerto %d esta abierto\n", i)
		}
	*/
	flag.Parse()
	var wg sync.WaitGroup

	for i := 0; i < 65535; i++ {
		go func(port int) {
			wg.Add(1)
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("El puerto %d esta abierto\n", port)
		}(i)
	}
	wg.Wait()
}
