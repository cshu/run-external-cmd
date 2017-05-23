package main

import (
	"os/exec"
	"flag"
	"encoding/hex"
)

var hexfl = flag.Bool("hex", false, "Decode args as base16")

func main() {
	flag.Parse()
	if flag.NArg()==0 {panic(1)}//os.exit(1) is okay but you don't write that everywhere
	sbuf:=flag.Args()
	if *hexfl{
		for i,v:=range sbuf{
			decres,err:=hex.DecodeString(v)
			if err!=nil{
				panic(1)
			}
			sbuf[i]=string(decres)
		}
	}
	cmd := exec.Command(sbuf[0], sbuf[1:]...)
	err:=cmd.Start()
	if err!=nil{
		panic(1)
	}
}
