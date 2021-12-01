package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//// 只读方式打开当前目录下的main.go文件
func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()
	//为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句

	//使用read方法读取数据
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed!, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))

	//io/ioutil，ioutil.ReadFile读取整个文件
	contents, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed!, err:", err)
		return
	}
	fmt.Println(string(contents))

	//bufio按行读取示例
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文明读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed err:", err)
			return
		}
		fmt.Print(line)
	}

	str := "hello 山河"
	file.Write([]byte(str)) //写入字节前切片数据
	file.WriteString("hello 小王子")
	//也可以用bufio.NewWriterr 和 ioutil.WriteFile写入文件
}
