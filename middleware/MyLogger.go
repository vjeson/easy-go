package middleware

import (
	"demo/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"time"
)

type Formatter  struct {

}



func logger(src io.Writer) gin.HandlerFunc{
	logger := logrus.New()

	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)


	//logger.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//})

	logger.SetFormatter(&logrus.JSONFormatter{

		TimestampFormat:"2006-01-02 15:04:05",

	})


	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//// 日志格式
		//logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//	statusCode,
		//	latencyTime,
		//	clientIP,
		//	reqMethod,
		//	reqUri,
		//)

		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}

func LoggerToFile() gin.HandlerFunc {

	logger := util.LogFile()

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		header := c.Request.Header
		proto := c.Request.Proto

		// 状态码
		statusCode := c.Writer.Status()


		// 请求IP
		clientIP := c.ClientIP()

		err := c.Err()

		body, _ := ioutil.ReadAll(c.Request.Body)


		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":     		err,
			"body": 		body,
		}).Info()
	}

}

func LoggerToStd() gin.HandlerFunc {
	logger := util.Log()

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//// 日志格式
		//logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//	statusCode,
		//	latencyTime,
		//	clientIP,
		//	reqMethod,
		//	reqUri,
		//)

		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}

