package main

import (
	zmq "github.com/pebbe/zmq4"
	"flag"
	"os"
	"fmt"
)

func main() {
	block := flag.Bool("block", true, "Use a blocking zmq_recv.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s <zmq_endpoint> [opts]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	uri := flag.Arg(0)
	sock, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create ZMQ socket: %v", err)
		os.Exit(2)
	}
	err = sock.Connect(uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not connect to endpoint %s: %v", uri, err)
		os.Exit(2)
	}
	err = sock.SetSubscribe("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not subscribe: %v", err)
		os.Exit(2)
	}
	var recvFlag zmq.Flag
	if *block {
		recvFlag = 0
	} else {
		recvFlag = zmq.DONTWAIT
	}
	for {
		msg, err := sock.RecvBytes(recvFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not receive bytes: %v", err)
			os.Exit(2)
		}
		os.Stdout.Write(msg)
	}
}
