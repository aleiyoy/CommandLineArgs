package main

import (
	"flag"
	"log"
)

func no1(){
	var name string
	flag.StringVar(&name, "name", "GO 语言变成之旅", "帮助信息")
	flag.StringVar(&name, "n", "GO 语言变成之旅", "帮助信息")

	flag.Parse()

	log.Printf("name1: %s",  name)
}

func no2(){
	var name string
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0{
		return
	}

	switch args[0] {
	case "go":
		// 返回带有指定名称的空命令集, 第二个参数是产生错误直接退出
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "GO 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		goCmd := flag.NewFlagSet("php", flag.PanicOnError)
		goCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	}

	log.Printf("name:%s", name)

}



func main(){
	//no1()
	no2()

}
