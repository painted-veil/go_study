package main

//一个TCP客户端进行TCP通信的流程如下：
/*
建立与服务端的链接
进行数据收发
关闭链接*/
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close() //关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { //输入Q就推出
			return
		}
		buf := [512]byte{}
		n, err := inputReader.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return

		}
		fmt.Println(string(buf[:n]))
	}
	//注意可能会出现粘包要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。因此要对数据包进行封包和拆包的操作
	/*由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，
	TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
	接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，
	然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。
	封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。包头部分的长度是固定的，
	并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。*/
}
