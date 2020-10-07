package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ProtocolICMP = 1
)

func Ping(addr string, seq int) (error) {
  c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
  if err != nil {
    return err
  }
  defer c.Close()

	dst, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		panic(err)
		return err
	}

	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: seq,
			Data: []byte("Hello!"),
		},
	}

	b, err := m.Marshal(nil)
	if err != nil {
		return err
	}

	_, err = c.WriteTo(b, dst)
	if err != nil {
		return err
	}
  return nil
}

func main() {
  seq := 1
	for
	{
      time.Sleep(1 * time.Second)
      fmt.Println("Ping!")
     
      err := Ping(os.Args[1], seq)
      if err != nil {
        fmt.Println("%v", err)
      }
      seq++
	 }
}
