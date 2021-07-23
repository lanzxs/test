package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"time"
)

func handleFunc() mgin.LimitReachedHandler {
	return func(c *gin.Context) {
		c.JSON(500,gin.H{"errmsg" : "xian su"})
		c.Abort()
		return
	}
}

func RateLimiter() gin.HandlerFunc {
	// 创建速率配置
	rate := limiter.Rate{
		Period: time.Second,
		Limit:  10,
	}
	// 将数据存入内存
	store := memory.NewStore()

	// 创建速率实例, 必须是真实的请求
	instance := limiter.New(store, rate, limiter.WithTrustForwardHeader(true))

	//global.Log.Infof("初始化速率限制中间件完成")

	// 生成gin中间件
	return mgin.NewMiddleware(instance,mgin.WithLimitReachedHandler(handleFunc()))
}
