package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

func worker(host string, ports, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		addr := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	host := flag.String("host", "", "host to scan")
	startPort := flag.Int("start", 1, "the start port to scan")
	endPort := flag.Int("end", 1024, "host to scan")
	workersCount := flag.Int("workers", 100, "number of workers")
	flag.Parse()

	if *host == "" {
		fmt.Println("Error: Please provide a host to scan.")
		return
	}

	ports := make(chan int, 2*(*workersCount))
	results := make(chan int, *workersCount)
	var openPorts []int
	var wg sync.WaitGroup

	for i := 0; i < *workersCount; i++ {
		wg.Add(1)
		go worker(*host, ports, results, &wg)
	}

	go func() {
		for p := *startPort; p < *endPort; p++ {
			ports <- p
		}
		close(ports)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	sort.Ints(openPorts)
	for _, p := range openPorts {
		fmt.Printf("%d open\n", p)
	}
}
