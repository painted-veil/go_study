package main

//ogger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
//Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
import (
	"fmt"
	"log"
	"os"
)

//如果要标准的Logger，则写道init函数中去
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile) //设置标准错误输出地，默认是标准错误输出
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {

	log.SetFlags(log.Llongfile | log.Lmsgprefix | log.Ldate)
	log.Println("这是一条日志")
	log.SetPrefix("[小王子]") //设置输出前缀
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")

	//og标准库中还提供了一个创建新logger对象的构造函数–New
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
