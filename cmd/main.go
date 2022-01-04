package main

import (
	"flag"
	"gitea.programmerfamily.com/go/product/internal/server"
	"log"
)

var (
	cfgPath = flag.String("c", "./configs/product.yaml", "service config path")
)

// 入口文件
func main()  {
	flag.Parse()

	if err := server.Init(*cfgPath); err != nil {
		log.Fatalf("config err:%+v",err)
	}

	server.Run()
}
