package main

import (
	"github.com/20100204/gotravel/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.execute err:%v",err)
	}
}
