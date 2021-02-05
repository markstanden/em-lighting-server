package main

import "net"

	 func main() {
	 	resolver := "udp"
		serverAddr := "192.168.1.1:16384"
		maxBufferSize := 123456

	 // init
    udpAddr, err := net.ResolveUDPAddr(resolver, serverAddr)
    if err != nil {
        // handle error
    }
    
    conn, err := net.ListenUDP("udp", udpAddr)
    if err != nil {
        // handle error
    }
  
    // send message
    buffer := make([]byte, maxBufferSize)
    n, addr, err := conn.ReadFromUDP(buffer)
    if err != nil {
        // handle error
    }

    // receive message
    buffer = make([]byte, maxBufferSize)
    n, err = conn.WriteToUDP(buffer[:n], addr)
    if err != nil {
        // handle error
    }
	}
/* /* package main

import (
	"fmt"
	"log"
	"net"
)

// Standard port that they use
const port string = "16384"

func main() {
	// listen to incoming udp packets
pc, err := net.ListenPacket("udp", ":16384")
if err != nil {
	log.Fatal(err)
}
defer pc.Close()

//simple read
buffer := make([]byte, 1024)
pc.ReadFrom(buffer)

//simple write
pc.WriteTo([]byte("Hello from client"), addr)
fmt.Println(buffer)
}

package main

import (
	"net"

	"fmt"
)

func main() {
  addr, _ := net.ResolveUDPAddr("udp", ":16384")
  sock, _ := net.ListenUDP("udp", addr)

  i := 0
  for {
    i++
    buf := make([]byte, 1024)
    rlen, _, err := sock.ReadFromUDP(buf)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(string(buf[0:rlen]))
    fmt.Println(i)
    //go handlePacket(buf, rlen)
  }
} */