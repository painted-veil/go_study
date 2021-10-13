package main

//io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
import (
	"fmt"
	"io"
	"strings"
)

//Read用数据填充给定的字节切片并返回填充字节数和错误值
//在遇到数据流的结尾时,会返回一个Io.EOF错误
//这里创建一个reader并以8字节速度读取输出
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
