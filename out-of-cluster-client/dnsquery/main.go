package main

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

func main() {
	fmt.Println("test dns query")
	g := new(dns.Msg)
	g.SetQuestion("hostsz.test-gslb.example.com.", dns.TypeA)

	a, _ := dns.Exchange(g, "127.0.0.1:53")
	for _, A := range a.Answer {
		IP := strings.Split(A.String(), "\t")[4]
		fmt.Println(IP)
	}
}
