package brute

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"h12.io/socks"
)

var T Transport

type Transport struct {
	Tr *http.Transport
}

func NewTor() Transport {
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
	T = Transport{&http.Transport{Dial: dialSocksProxy}}
	return T
}

func (transport *Transport) Check_IP() {
	var body []byte
	httpClient := &http.Client{Transport: T.Tr}
	resp, err := httpClient.Get("http://bot.whatismyipaddress.com/")
	body, err = ioutil.ReadAll(resp.Body)
	check_error(err)
	fmt.Println("IP de requests:" + string(body))
	resp.Body.Close()

}
