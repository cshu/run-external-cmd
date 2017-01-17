package main

import (
	"os"
	"os/exec"
	"flag"
	"encoding/hex"
)

var hexfl = flag.Bool("hex", false, "Decode args as base16")

func main() {
	flag.Parse()
	if flag.NArg()==0 {os.Exit(1)}
	sbuf:=flag.Args()
	if *hexfl{
		for i,v:=range sbuf{
			decres,err:=hex.DecodeString(v)
			sbuf[i]=string(decres)
			if err!=nil{
				os.Exit(1)
			}
		}
	}
	cmd := exec.Command(sbuf[0], sbuf[1:]...)
	err:=cmd.Start()
	if err!=nil{
		os.Exit(1)
	}
}
