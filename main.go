package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	count := flag.Int("c", 100, "Count")
	interval := flag.Int("i", 1, "Interval in seconds")
	flag.Parse()
	scheme, host, port, path := parseUrl(flag.Arg(0))
	ip := ipLookup(host)

	for i := 1; ; i++ {
		response, elapsed := httpGet(scheme + "://" + host + ":" + port + path)
		fmt.Printf("connected to %s:%s, seq=%d, time=%d ms, response=%s\n", ip, port, i, elapsed, response)
		if i >= *count {
			break
		} else {
			time.Sleep(time.Duration(*interval) * time.Second)
		}
	}
}

func parseUrl(s string) (string, string, string, string) {
	var host string
	var port string

	if !strings.HasPrefix(s, "http") {
		s = "http://" + s
	}

	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	scheme := u.Scheme
	path := u.Path

	if strings.Contains(u.Host, ":") {
		host, port, _ = net.SplitHostPort(u.Host)
	} else {
		host = u.Host
		if strings.Contains(u.Scheme, "https") {
			port = "443"
		} else {
			port = "80"
		}
	}

	return scheme, host, port, path
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

func ipLookup(s string) string {
	addr, err := net.LookupHost(s)
	if err != nil {
		log.Fatal(err)
	}
	return addr[0]
}
