package router

import (
	"github.com/cuiyuanxin/airuisi-admin/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"say": "Hi airuisi admin!",
		})
	})
	// 测试通信
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "pong",
		})
	})
	//// 限流器
	//var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	//	Key:          "/auth",
	//	FillInterval: time.Second,
	//	Capacity:     10,
	//	Quantum:      10,
	//})
	//
	//// 链路追踪
	//r.Use(middleware.Tarcing())
	//if global.ServerSetting.GinMode == "debug" {
	//	r.Use(gin.Logger())
	//	r.Use(gin.Recovery())
	//} else {
	//	r.Use(middleware.AccessLog())
	//	r.Use(middleware.Recovery())
	//}
	//
	//r.Use(middleware.RateLimiter(methodLimiters))
	//r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout * time.Second))
	r.Use(middleware.Cors())
	//
	////r.StaticFS("/static/qrcode", http.Dir(global.WxInfoSetting.QrcodeSavePath))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	//r = route.Router(r)
	//
	return r
}
