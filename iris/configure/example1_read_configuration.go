package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Appname string	`json:"appname"`
	Port int	`json:"port"`
}

func main() {
	file, _ := os.Open("E:\\chengzhen\\cz-goland-workspqce\\src\\GoLearning\\iris\\configure\\config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Configuration{}

	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(conf.Port)
}
/*
细节点：
1 结构体的成员首字母大写(结构体成员首字母小写会读不出来，config.json里面字段大小写都行)。小写是内部私有，大写外部才能访问
2 配置文件的配置项须与结构体成员名一样
3 支持bool， int， float ， 字符串，字符串数组...等，也可以包含其他结构体 如[Friend]
*/