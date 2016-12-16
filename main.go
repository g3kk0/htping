package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func targetParser(s string) (string, string, string){
	//protocol
	protoExpr, err := regexp.Compile("^https?")
	if err != nil {
		log.Fatal(err)
	}

	var protocol string
	if protoExpr.MatchString(s) {
		protocol = protoExpr.FindString(s)
	} else {
		protocol = "http"
	}

	// port
	portExpr, err := regexp.Compile(":\\d*$")
	if err != nil {
		log.Fatal(err)
	}

	var port string
	if portExpr.MatchString(s) {
		port = portExpr.FindString(s)
		port = port[1:]
	} else if protocol == "https" {
		port = "443"
	} else {
		port = "80"
	}

	//domain
	domainExpr, err := regexp.Compile("^https?://")
	if err != nil {
		log.Fatal(err)
	}

	domain := domainExpr.ReplaceAllString(s, "")
	domain = portExpr.ReplaceAllString(domain, "")

	return protocol, domain, port
}

func ip(s string) string {
	addr, err := net.LookupHost(s)
	if err != nil {
		log.Fatal(err)
	}
	return addr[0]
}

func httpGet(s string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(s)
	if err != nil {
		log.Fatal(err)
	}

	statusCode := strconv.Itoa(resp.StatusCode)

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "\033[32m" + statusCode + "\033[0m"
	} else if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		return "\033[33m" + statusCode + "\033[0m"
	} else {
		return "\033[31m" + statusCode + "\033[0m"
	}
}

func main() {
	target := os.Args[1]
	protocol, domain, port := targetParser(target)

	for i := 0; ; i++ {
		fmt.Printf("connected to %s:%s, seq=%d, time=0 ms, response=%s\n", ip(domain), port, i, httpGet(protocol+"://"+domain+":"+port))
		time.Sleep(1000 * time.Millisecond)
	}
}
