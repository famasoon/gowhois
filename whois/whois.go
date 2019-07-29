package whois

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"github.com/famasoon/gowhois/server"

	"golang.org/x/net/publicsuffix"
)

const WHOIS_PORT = "43"

func Whois(domain string) (result string, err error) {
	domain = strings.Trim(strings.TrimSpace(domain), ".")
	if domain == "" {
		err = fmt.Errorf("Domain is Empty")
		return
	}

	fmt.Println(domain)

	result, err = whoisQuery(domain)
	if err != nil {
		return
	}

	return result, nil
}

func whoisQuery(domain string) (result string, err error) {
	var whoisServer string

	if !isIpv4(domain) {
		eTld, _ := publicsuffix.PublicSuffix(domain)
		eTld = "." + eTld
		whoisServer = server.GetWhoisServer(eTld)
	} else {
		whoisServer = server.GetWhoisServer(domain)
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(whoisServer, WHOIS_PORT), time.Second*30)
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))

	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return
	}

	result = string(buffer)

	return result, nil
}

func isIpv4(s string) bool {
	result := net.ParseIP(s)
	return result != nil
}
