// package zscan
package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)


func worker(host string, startPort, endPort int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := startPort; p < endPort; p++ {
		addr := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		func (){if err := conn.Close; err != nil {
			return 
		}}()
		results <- p
	}
}

func main() {
	host := flag.String("host", "", "host to scan")
	startPort := flag.Int("start", 1, "the start port to scan")
	endPort := flag.Int("end", 1024, "host to scan")
	workersCount := flag.Int("workers", 10, "number of workers")
	flag.Parse()

	if *host == "" {
		fmt.Println("Error: Please provide a host to scan.")
		return
	}

	portsPerWorker := (*endPort - *startPort) / *workersCount
	remainingPorts := (*endPort - *startPort) % *workersCount

	var openPorts []int
	var wg sync.WaitGroup
	results := make(chan int, *endPort-*startPort)

	for i := 0; i < *workersCount; i++ {
		wg.Add(1)
		start := *startPort + i*portsPerWorker
		end := start + portsPerWorker
		if i == *workersCount-1 {
			end += remainingPorts
		}

		go worker(*host, start, end, results, &wg)
	}

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
