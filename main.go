package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	count := flag.Int("c", 100, "Count")
	interval := flag.Int("i", 1, "Interval")
	flag.Parse()
	target := flag.Arg(0)
	protocol, domain, port := targetParser(target)
	ip := ip(domain)

	for i := 0; ; i++ {
		if i >= *count {
			break
		}
		response, elapsed := httpGet(protocol + "://" + domain + ":" + port)
		fmt.Printf("connected to %s:%s, seq=%d, time=%d ms, response=%s\n", ip, port, i, elapsed, response)
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

func httpGet(s string) (string, int64) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	start := time.Now()
	resp, err := client.Get(s)
	if err != nil {
		return "\033[31mNo response\033[0m", 0
	}
	elapsed := int64(time.Since(start) / time.Millisecond)

	statusCode := strconv.Itoa(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "\033[32m" + statusCode + "\033[0m", elapsed
	} else if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		return "\033[33m" + statusCode + "\033[0m", elapsed
	} else {
		return "\033[31m" + statusCode + "\033[0m", elapsed
	}
}

func ip(s string) string {
	addr, err := net.LookupHost(s)
	if err != nil {
		log.Fatal(err)
	}
	return addr[0]
}

func targetParser(s string) (string, string, string) {
	// protocol
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

	// domain
	domainExpr, err := regexp.Compile("^https?://")
	if err != nil {
		log.Fatal(err)
	}

	domain := domainExpr.ReplaceAllString(s, "")
	domain = portExpr.ReplaceAllString(domain, "")

	return protocol, domain, port
}
