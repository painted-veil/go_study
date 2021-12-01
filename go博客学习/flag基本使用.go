//Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	//flag.Type(flag名, 默认值, 帮助信息)*Type
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
	//需要通过调用flag.Parse()来对命令行参数进行解析。

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)

	//返回命令行其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	flag.NArg()
	//返回使用的命令行参数个数
	flag.NFlag()
}
