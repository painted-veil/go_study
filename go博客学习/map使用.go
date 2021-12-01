//Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现
//map是一种无序的基于key-value数据结构，go语言中的map是引用类型必须初始化才能使用
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["牛鼓励"] = 60
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123454",
	}
	fmt.Println(userInfo)

	//判断某个键是否存在
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
	//map的遍历，使用for range遍历map
	//与添加键值对的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//使用delete函数删除键值对
	delete(scoreMap, "小明")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	//按照指定顺序遍历map

	//元素map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	//初始化操作
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	//map中值为切片类型的操作
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	//初始化操作
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}
