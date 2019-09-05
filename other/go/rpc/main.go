package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Args rpc args struct
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Mutiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Hello(name string, reply *string) error {
	*reply = "Hello, " + name
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":11223")
	if e != nil {
		log.Fatal(e)
	}

	go http.Serve(l, nil)

	client, e := rpc.DialHTTP("tcp", "localhost:11223")
	if e != nil {
		log.Fatal(e)
	}

	args := &Args{4, 3}
	var reply int
	e = client.Call("Arith.Mutiply", args, &reply)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("Result: %d\n", reply)

	var hello string
	e = client.Call("Arith.Hello", "World", &hello)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("Hello Reply: %s\n", hello)
}
