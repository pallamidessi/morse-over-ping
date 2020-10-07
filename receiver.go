package main

import (
	"fmt"
	"time"
	"golang.org/x/net/icmp"
)

const (
	ProtocolICMP = 1
)

func WaitForPing() (time.Duration, error) {
  c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
  if err != nil {
    return 0, err
  }
  defer c.Close()

  reply := make([]byte, 1500)
  
  n, _, err := c.ReadFrom(reply)
  if err != nil {
    return 0, err
  }

  rm, err := icmp.ParseMessage(ProtocolICMP, reply[:n])
  if err != nil {
    return 0, err
  }

  fmt.Println("Got %v!", rm)
  return 0, nil
}

func main() {
  seq := 1
	for
	{
      time.Sleep(1 * time.Second)
      fmt.Println("Pong!")

      _, err := WaitForPing()
      if err != nil {
        fmt.Println("%v", err)
      }
      seq++
	 }
}
