package main

import (
	"flag"
	"fmt"
	"Trovilo/brute"
)

func findDashboard(wordlist []string, machine brute.Conf, url string) {
	for _, endswith := range wordlist {
		res := machine.Request(url + endswith)
		if res["status_code"] == "200" {
			brute.Write(res["url"])
		}
	}
}

func parse(Wordlist string, Tor bool, Url string) {
	var conf brute.Conf
	param := brute.Parser(Wordlist)
	if Tor {
		tor := brute.NewTor()
		tor.Check_IP()
		conf.Set_Tor(true)
	}
	for _, wordlist := range param {
		conf := brute.NewConf(wordlist)
		go findDashboard(conf.Content, conf, Url)
	}
}

func main() {
	brute.Banner()
	var Wordlist = flag.String("Wordlist", "0", "Wordlist Path.")
	var Url = flag.String("url", "http://pudim.com.br/", "Website URL")
	var Tor = flag.Bool("tor", false, "Set up a tor proxy")
	flag.Parse()
	parse(*Wordlist, *Tor, *Url)
	var input string
	fmt.Scanln(&input)
}
