package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
测试运行：
1 go run main.og
2 curl "localhost:8085" 请求
3 Ctrl + C 停止进程
4 curl 会处理请求 返回Welcome Gin Server
*/

/**
1 响应http请求的部分放入goruntine中
2 通过信道阻塞主线程，监听信号
3 接收到退出信号时，做清理工作
4 退出
 */

//优雅关停服务  当我们部署的服务被关闭时，如果正在运行一些操作（读写数据等），则可能会导致数据的丢失。通常我们希望服务被关闭之前，能处理好正在进行的操作之后再退出
/*
传统关停服务器的方式，开始有一个gin实例，gin.run通过阻塞监听端口，请求过来之后会请求回调函数提供服务，结束之后就直接终止了
优雅关停的方式，可以使用一个server.ListenAndServer 构建一个server来代替gin.Run，它是不阻塞的。另外使用os.Signal来阻塞进程，监听关闭信号，如果获取到信号就将超时的上下文传递到server的shutdwon方法里才正式退出
*/
func main() {
	router := gin.Default()
	//设置路由
	router.GET("/", func(c *gin.Context) {
		//模拟超时请求
		time.Sleep(5 * time.Second)
		fmt.Println("Welcome Gin Server11111")
		c.String(http.StatusOK, "Welcome Gin Server22222")
	})

	srv := &http.Server{
		Addr:    ":8085",
		Handler: router,
	}
	//开启一个协程用来监听关闭信号
	//响应http请求的部分放入goruntine中
	go func() {
		//服务连接
		//ListenAndServe 不阻塞
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	//通过信道阻塞主线程，监听信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Shutdown Server by sig: %v", sig)
	log.Println("Shutdown Server ...")
	// 接收到退出信号时，做清理工作
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
