// 服务器端
package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

type Common struct {
	Age  int
	Name string
}

type Mystring string

func (s *Mystring) Name(common *Common, desc *string) error {
	*desc = fmt.Sprintf("名字是： %s,年龄是：%d", common.Name, common.Age)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	comm := new(Mystring)
	rpc.Register(comm)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	os.Stdin.Read(make([]byte, 1))
}
