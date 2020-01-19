package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	fmt.Println("test dns query")
	g := new(dns.Msg)
	g.SetQuestion("app3.cloud.example.com.", dns.TypeA)

	a, _ := dns.Exchange(g, "127.0.0.1:53")
	fmt.Println(a.Answer)
}
