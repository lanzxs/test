// 客户端
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

/*type Args struct {
    A, B int
}

type Common struct {
    Age int
    Name string
}*/

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//然后，客户端可以执行远程调用：

	//同步调用
	comm := &Common{24, "lan"}
	var desc string
	/*err = client.Call("Mystring.Name",comm,&desc)
	if err != nil {
		log.Fatal("错误了",err)
	}

	fmt.Println("result:",desc)*/

	//异步调用
	syncChan := client.Go("Mystring.Name", comm, &desc, nil)
	for {
		select {
		case <-syncChan.Done:
			fmt.Println("result:", desc)
			return
		default:
			fmt.Println("等待结果返回")
			time.Sleep(time.Second * 1)
		}
	}
	// Synchronous call
	/*args := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
	    log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)*/
}
