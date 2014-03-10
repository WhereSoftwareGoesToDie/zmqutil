package main

import (
	zmq "github.com/pebbe/zmq4"
	"flag"
	"os"
	"fmt"
)

func Exitf(code int, format string, v... interface{}) {
	fmt.Fprintf(os.Stderr, "Fatal: ")
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(code)
}

func main() {
	block := flag.Bool("block", true, "Use a blocking zmq_recv.")
	text := flag.Bool("text", false, "Read messages as strings.")
	multipart := flag.Bool("multipart", false, "Read multipart messages and print the parts individually.")
	flag.Usage = func() {
		Exitf(2, "Usage of %s:\n", os.Args[0])
		Exitf(2, "%s <zmq_endpoint> [opts]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}
	uri := flag.Arg(0)
	sock, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		Exitf(2, "Could not create ZMQ socket: %v", err)
	}
	err = sock.Connect(uri)
	if err != nil {
		Exitf(2, "Could not connect to endpoint %s: %v", uri, err)
	}
	err = sock.SetSubscribe("")
	if err != nil {
		Exitf(2, "Could not subscribe: %v", err)
	}
	var recvFlag zmq.Flag
	if *block {
		recvFlag = 0
	} else {
		recvFlag = zmq.DONTWAIT
	}
	for {
		if *multipart && *text {
			msg, err := sock.RecvMessage(recvFlag)
			if err != nil {
				Exitf(2, "Could not receive bytes: %v", err)
			}
			for _, part := range msg {
				fmt.Println(part)
			}
		} else if *multipart {
			msg, err := sock.RecvMessageBytes(recvFlag)
			if err != nil {
				Exitf(2, "Could not receive bytes: %v", err)
			}
			for _, part := range msg {
				os.Stdout.Write(part)
			}
		} else if *text {
			msg, err := sock.Recv(recvFlag)
			if err != nil {
				Exitf(2, "Could not receive bytes: %v", err)
			}
			fmt.Println(msg)
		}
	}
}
