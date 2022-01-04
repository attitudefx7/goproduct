package server

import (
	"gitea.programmerfamily.com/go/product/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 全局捕获异常处理
func Recover(c *gin.Context) {
	c.Set("start_time", time.Now())

	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			logger.GetLogger().Errorf("panic: %v", r)
			log.Printf("panic: %v\n", r)
			//debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			c.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()

	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func LogRecord(c *gin.Context)  {

	defer func() {

		logs := make(map[string]interface{})
		logs["url"] = c.Request.URL.Path
		logs["ip"] = c.ClientIP()
		logs["params"] = c.Params
		logs["message"] = c.Request.Response

		startTime, ok := c.Get("start_time")
		if ok {
			switch startTime.(type) {
			case time.Time:
				logs["use_time"] = time.Now().Sub(startTime.(time.Time))
			}
		}


		logger.GetLogger().Debug(logs)
	}()

	//加载完 defer recover，继续后续接口调用
	c.Next()
}