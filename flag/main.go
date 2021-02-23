package main

import (
	"flag"
	"fmt"
)

func main() {
/*	var name string
	flag.StringVar(&name,"name","go 语言编程之旅","帮助信息")
	flag.StringVar(&name,"n","go go go","help help")
	flag.Parse()
	log.Printf("name:%s",name)
	fmt.Println()*/
	var secname string
	flag.Parse()
	goCmd := flag.NewFlagSet("go",flag.ExitOnError)
	goCmd.StringVar(&secname,"name","go language","help golanguage")
	phpCmd := flag.NewFlagSet("php",flag.ExitOnError)
	phpCmd.StringVar(&secname,"n","php language","help php")
	args := flag.Args()
fmt.Println(args)
	switch args[0] {
	case "go":
		_=goCmd.Parse(args[1:])
	case "php":
		_=phpCmd.Parse(args[1:])
	}
	fmt.Printf("secname:%s",secname)
}
