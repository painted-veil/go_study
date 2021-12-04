package main

import (
	"bufio"
	"fmt"
	"net"
)

//用go实现Tcp客户端
//TCP服务端程序处理流程
/**监听端口
接收客户端请求建立链接
创建goroutine处理链接。*/
func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("readn from cient failed,err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到clinet发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) //发送数据

	}
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("listen faileed error: ", err)
		return
	}
	for {
		conn, err := listener.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}
