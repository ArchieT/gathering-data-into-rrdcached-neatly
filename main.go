package main

import (
	"github.com/multiplay/go-rrd"
	"os"
)

func main() {
	c := rrd.NewClient("/tmp/rrdcached.sock",rrd.Unix)
	argsWOProg := os.Args[1:]
}
