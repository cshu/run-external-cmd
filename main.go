package main

import (
	"encoding/hex"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

var hexfl = flag.Bool("hex", false, "Decode args as base16.")
var waitstdinfl = flag.Bool("wait-stdin", false, "Do not exit immediately. Wait until stdin EOF. Also kill child process before exit.")
var coutfl = flag.Bool("cout", false, "Pipe stdout.")
var cerrfl = flag.Bool("cerr", false, "Pipe stderr.")
var chdirfl = flag.String("C", "", "Change the working directory to dir before invoking command (similar to git/env).")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		panic(1)
	}
	sbuf := flag.Args()
	if *hexfl {
		for i, v := range sbuf {
			decres, err := hex.DecodeString(v)
			if err != nil {
				panic(err)
			}
			sbuf[i] = string(decres)
		}
	}
	cmd := exec.Command(sbuf[0], sbuf[1:]...)
	if *chdirfl != "" {
		cmd.Dir = *chdirfl
	}
	var stdout io.ReadCloser
	var stderr io.ReadCloser
	var err error
	if *coutfl {
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
	}
	if *cerrfl {
		stderr, err = cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	if *coutfl {
		wg.Add(1)
		go func(stdstream io.ReadCloser) {
			defer wg.Done()
			_, err := io.Copy(os.Stdout, stdstream)
			if err != nil {
				panic(err)
			}
		}(stdout)
	}
	if *cerrfl {
		wg.Add(1)
		go func(stdstream io.ReadCloser) {
			defer wg.Done()
			_, err := io.Copy(os.Stderr, stdstream)
			if err != nil {
				panic(err)
			}
		}(stderr)
	}
	if *waitstdinfl {
		go func() {
			_, err = io.Copy(ioutil.Discard, os.Stdin)
			cmd.Process.Kill()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}()
	}
	if *waitstdinfl || *coutfl || *cerrfl {
		wg.Wait()         //note you cannot remove this line and simply use cmd.Wait() to wait for EOF, because the goroutine doing stdout/stderr reading might have not started running yet. So you have a slim chance getting error "file already closed"
		err := cmd.Wait() //>Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}
