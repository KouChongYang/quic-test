package main

import (
	"flag"
	"io/ioutil"
	"fmt"
)

const DEFAULT_ADDR = "0.0.0.0:4443"

func usage() {
	fmt.Println("Usage: quictest [-s [host] -F seed_file |-c host] [options]")
}

func main() {
	
	c:=""
	n:=0
	addr:=""
	file:=""
	ticks:=0
	cn:=0
	flag.StringVar(&c, "c", "127.0.0.1:34567", "run client connect to server")
	flag.IntVar(&n, "n", 1, "number of streams")
	flag.IntVar(&cn, "cn", 1, "number of connect")
	flag.StringVar(&addr, "b", "0.0.0.0:34567", "bind to")
	flag.StringVar(&file, "f", "./a", "source file")
	s := flag.Bool("s", false, "Run server")
	flag.IntVar(&ticks,"t", 10, "Time in seconds to transmit for (default 10 secs)")
	
	flag.Parse()
	fmt.Println("=========cn:",cn)	
	fmt.Println("s:",*s,"c:",c,"stream:",n,"addr:",addr,"file:",file,"ticks:",ticks)

	if *s {
		fmt.Println("================")
		if file == "" {
			usage()
			return
		}
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting server...")
		go serverMain(addr, data, n, ticks)
		serverMain("0.0.0.0:34567", data, n, ticks)
	} else {
		if c == "" {
			usage()
			return
		}
		fmt.Println("Starting client...")
		for i:= 0;i<cn;i++{
	go			clientMain(c)
		}
			clientMain(c)
	}
}
