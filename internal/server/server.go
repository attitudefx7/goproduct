package server

import (
	"context"
	"fmt"
	"gitea.programmerfamily.com/go/product/internal/db"
	"gitea.programmerfamily.com/go/product/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var config *Config

func Init(cfgPath string) error {
	var err error
	// 通过path加载配置项
	config, err = newConfigFromFile(cfgPath)
	if err != nil {
		return err
	}

	// 加载DB配置
	if err = db.Init(config.Db); err != nil {
		return err
	}

	// 加载logger配置
	logger.Init(config.Logger)

	return nil
}

func Run()  {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	// 初始化
	engine := gin.Default()

	// 注册路由
	registerRoute(engine)

	// ---------------------------------------------------------------
	// + 优雅的关闭服务
	// + 优雅就是不暴力关闭服务器，等待处理完毕再关闭
	// + 使用 go 协程 启动服务
	// + 为了防止长时间处理为完毕，所以需要设置超时时间
	// + http.Server 内置的 Shutdown()方法支持优雅地关机
	// + 所以要使用 http.Server
	// ---------------------------------------------------------------

	// 设置srv参数
	srv := http.Server{
		Handler: engine,
		Addr: fmt.Sprintf(":%d",config.Port),
	}

	go func() {
		// 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	// notify 不会阻塞
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里设置阻塞，接收到信号后才会往下执行
	<- quit
	log.Println("Shutdown Server ...")

	// 设置一个带超时的ctx
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// 关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %v", err)
	}

	log.Println("Server exiting")
}
